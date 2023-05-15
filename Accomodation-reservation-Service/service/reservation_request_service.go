package service

import (
	"Accomodation-reservation-Service/communication"
	"Accomodation-reservation-Service/domain"
	"Accomodation-reservation-Service/repo"
	"context"
	"fmt"
	pb "github.com/XWS-tim24/Common/common/proto/accommodation_service"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type ReservationRequestService struct { //Accept prihvata rezervaciju, updatuje u bazi i poziva createReservation
	ReservationRequestRepo     *repo.ReservationRequestRepository
	ReservationService         *ReservationService
	AccommodationServiceAddres string
}

func (service *ReservationRequestService) GetAll() (*[]domain.ReservationRequest, error) {
	return service.ReservationRequestRepo.GetAll()
}

func (service *ReservationRequestService) GetById(id string) (*domain.ReservationRequest, error) {
	reservationRequest, err := service.ReservationRequestRepo.GetById(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("reservation request with id %s not found", id))
	}
	return &reservationRequest, nil
}

func (service *ReservationRequestService) Create(reservationRequest *domain.ReservationRequest) error {
	//done: 1. Proveri u svojoj bazi da li vec postoji REZERVACIJA za taj SMESTAJ u tom DATUMU sa Active(0) STATUSOM
	//2. Kontaktirati accomodation service da li je smestaj dostupan u tom periodu (vlasnik odredjuje dostupnost i to se cuva u acc servisu)
	//3. Kontaktirati accomodation service da li je rucna ili automatska potvrda:
	//	 	automatska -> poziva se funkcija Potvrdi rez (ne treba da proverava preklapajuce zahteve jer ih nema)
	//Ako nesto od ova 3 ne prodje status Denied

	if reservationRequest.StartDate.After(reservationRequest.EndDate) {
		return fmt.Errorf(fmt.Sprintf("Start date must be before end date, request id %s", reservationRequest.Id))
	}

	reservationRequest.Status = domain.Pending
	reservationRequest.Deleted = false
	if service.AlreadyReservedForDate(reservationRequest.AccomodationId, reservationRequest.StartDate, reservationRequest.EndDate) { //ili 2
		return fmt.Errorf("accommodation is already reserved for chosen date") //service.createDenied(reservationRequest)
	}
	accommodationClient := communication.NewAccommodationClient(service.AccommodationServiceAddres)
	fmt.Println("checking availability for :id", reservationRequest.AccomodationId)
	request := &pb.TimeSlotAvailableRequest{}
	request.AvailableTimeSlotDTO = &pb.AvailableTimeSlotDTO{AccommodationId: reservationRequest.AccomodationId, StartDate: timestamppb.New(reservationRequest.StartDate), EndDate: timestamppb.New(reservationRequest.EndDate)}
	response, err := accommodationClient.TimeSlotAvailableForAccommodation(context.TODO(), request)
	if err != nil {
		return err
	}
	if !response.Available {
		return fmt.Errorf("accommodation is not available in time slot")
	}

	fmt.Println("checking if automatic confirmation is chosen for accommodation id :", reservationRequest.AccomodationId)
	request1 := &pb.GetByIdRequest{Id: reservationRequest.AccomodationId}
	response1, err1 := accommodationClient.GetAutomaticAcceptById(context.TODO(), request1)
	if err1 != nil {
		return err1
	}
	if response1.AutomaticAccept {
		reservationRequest.Status = domain.Accepted
		err := service.ReservationRequestRepo.Create(reservationRequest)
		if err != nil {
			return err
		}
		fmt.Println("created accepted request with id :", reservationRequest.Id)
		err = service.createReservation(reservationRequest)
		if err != nil {
			return err
		}
		return nil
	}

	err = service.ReservationRequestRepo.Create(reservationRequest)
	if err != nil {
		return err
	}
	return nil
}

func (service *ReservationRequestService) Accept(id string) error {
	reservationRequest, _ := service.GetById(id)
	if reservationRequest.StartDate.Before(time.Now()) {
		return fmt.Errorf(fmt.Sprintf("reservation request with id %s has expired", id))
	}
	if reservationRequest.Status != domain.Pending {
		return fmt.Errorf(fmt.Sprintf("only pending requests are allowed to accept. request id %s", id))
	}

	if reservationRequest.Deleted {
		return fmt.Errorf(fmt.Sprintf(" request is deleted. request id %s", id))
	}

	reservationRequest.Status = domain.Accepted
	err := service.ReservationRequestRepo.AcceptOrDeny(reservationRequest)
	if err != nil {
		print("error during accepting request, request id %s", reservationRequest.Id.String())
		return err
	}
	err = service.createReservation(reservationRequest)
	if err != nil {
		print("error during accepting request, request id %s", reservationRequest.Id.String())
		return err
	}
	err = service.denyOthers(reservationRequest)
	if err != nil {
		print("error during accepting request, request id %s", reservationRequest.Id.String())
		return err
	}
	return nil
}

func (service *ReservationRequestService) Deny(id string) error {
	reservationRequest, _ := service.GetById(id)
	if reservationRequest.Deleted {
		return fmt.Errorf(fmt.Sprintf(" request is deleted. request id %s", id))
	}
	reservationRequest.Status = domain.Denied
	err := service.ReservationRequestRepo.AcceptOrDeny(reservationRequest)
	if err != nil {
		print("error during accepting request, request id %s", reservationRequest.Id.String())
		return err
	}
	return nil
}

func (service *ReservationRequestService) Delete(id string) error {
	err := service.ReservationRequestRepo.Delete(id)
	if err != nil {
		print("error during deleting reservation request logicaly, request id %s", id)
		return err
	}
	return nil
}

// METODE ZA FRONT
func (service *ReservationRequestService) GetAllPendingForUser(userId string) *[]domain.ReservationRequest {
	return service.ReservationRequestRepo.GetAllPendingForUser(userId)
}

func (service *ReservationRequestService) GetAllPendingForAccomodation(userId string) *[]domain.ReservationRequest {
	return service.ReservationRequestRepo.GetAllPendingForAccomodation(userId)
}

// POMOCNE METODE
func (service *ReservationRequestService) AlreadyReservedForDate(accomodationId string, startDate time.Time, endDate time.Time) bool {
	exists := service.ReservationRequestRepo.AlreadyReservedForDate(accomodationId, startDate, endDate)
	return exists
}

func (service *ReservationRequestService) createDenied(reservationRequest *domain.ReservationRequest) error {
	reservationRequest.Status = domain.Denied
	err := service.ReservationRequestRepo.Create(reservationRequest)
	return err
}

func (service *ReservationRequestService) createReservation(reservationRequest *domain.ReservationRequest) error {
	reservation := domain.Reservation{}
	reservation.RequestId = reservationRequest.Id.String()
	reservation.Status = domain.Active
	err := service.ReservationService.Create(&reservation)
	return err
}

func (service *ReservationRequestService) denyOthers(reservationRequest *domain.ReservationRequest) error {
	err := service.ReservationRequestRepo.DenyOthers(reservationRequest)
	return err
}
