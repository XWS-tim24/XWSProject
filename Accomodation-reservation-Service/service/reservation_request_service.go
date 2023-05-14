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
	if service.alreadyReservedForDate(reservationRequest) { //ili 2
		return service.createDenied(reservationRequest)
	}
	accommodationClient := communication.NewAccommodationClient(service.AccommodationServiceAddres)
	request := &pb.TimeSlotAvailableRequest{}
	request.AvailableTimeSlotDTO = &pb.AvailableTimeSlotDTO{AccommodationId: reservationRequest.AccomodationId, StartDate: timestamppb.New(reservationRequest.StartDate), EndDate: timestamppb.New(reservationRequest.EndDate)}
	response, err := accommodationClient.TimeSlotAvailableForAccommodation(context.TODO(), request)
	if err != nil {
		return err
	}
	if response.Available == false {
		return fmt.Errorf("accommodation is not available in time slot")

	}
	//automatska_potvrda := true //3.
	//if automatska_potvrda {
	//	reservationRequest.Status = domain.Accepted
	//service.MakeReservation(reservationRequest)
	//}
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
func (service *ReservationRequestService) alreadyReservedForDate(reservationRequest *domain.ReservationRequest) bool {
	exists := service.ReservationRequestRepo.AlreadyReservedForDate(reservationRequest)
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
