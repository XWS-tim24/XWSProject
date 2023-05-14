package handler

import (
	"Accomodation-Service/service"
	"context"
	"encoding/json"
	"fmt"
	"log"

	pb "github.com/XWS-tim24/Common/common/proto/accommodation_service"
)

type AccommodationHandler struct {
	pb.UnimplementedAccommodationServiceServer
	AccommodationService *service.AccommodationService
	AvailableDateService *service.AvailableDateService
}

func (handler *AccommodationHandler) CreateAcc(ctx context.Context, pbAcc *pb.CreateAccommodationRequest) (*pb.CreateAccommodationResponse, error) {
	accommodation := mapToAccommodation(pbAcc.Accommodation)
	println("Body:")

	err := handler.AccommodationService.Create(accommodation)
	if err != nil {
		return nil, err
	}
	return &pb.CreateAccommodationResponse{}, nil
}

func (handler *AccommodationHandler) GetAccommodationById(ctx context.Context, req *pb.GetByIdRequest) (*pb.CreateAccommodationResponse, error) {
	id := req.Id

	println("Accommodation with id %s", id)
	accommodation, err := handler.AccommodationService.GetById(id)

	if err != nil {
		return nil, err
	}
	pbAccommodation := mapToPbAccommodation(accommodation)
	return &pb.CreateAccommodationResponse{Accommodation: pbAccommodation}, nil
}

func (handler *AccommodationHandler) SearchAccommodations(ctx context.Context, req *pb.SearchRequest) (*pb.SearchResponse, error) {

	accommodationSearchDTO := mapAccommodationSearchDTO(req.AccommodationSearchDTO)
	println("Body:")
	accStr, _ := json.Marshal(accommodationSearchDTO)
	println(string(accStr))

	accommodations := handler.AccommodationService.Search(accommodationSearchDTO)
	response := &pb.SearchResponse{
		Accommodation: []*pb.Accommodation{},
	}
	for _, acc := range *accommodations {
		current := mapToPbAccommodation(&acc)
		response.Accommodation = append(response.Accommodation, current)
	}
	return response, nil
}

func (handler *AccommodationHandler) CreateAvailableDate(ctx context.Context, pbAvailableDate *pb.CreateAvailableDateRequest) (*pb.CreateAvailableDateResponse, error) {

	availableDate := mapToAvailableDate(pbAvailableDate.AvailableDate)

	println("Body:")
	availableDateJson, _ := json.Marshal(availableDate)
	println(string(availableDateJson))

	err := handler.AvailableDateService.Create(availableDate)
	if err != nil {
		println("Error while creating a new Available Date")
		return nil, err
	}
	return nil, nil
}

func (handler *AccommodationHandler) GetAvailableDateById(ctx context.Context, req *pb.GetByIdRequest) (*pb.CreateAvailableDateResponse, error) {
	id := req.Id
	log.Printf("AvailableDate with id %s", id)
	availableDate, err := handler.AvailableDateService.GetById(id)

	if err != nil {
		return nil, err
	}

	availableDatePb := mapToAvailableDatePb(availableDate)
	return &pb.CreateAvailableDateResponse{AvailableDate: availableDatePb}, nil
}

func (handler *AccommodationHandler) UpdateAvailableDate(ctx context.Context, req *pb.UpdateAvailableDateRequest) (*pb.UpdateAvailableDateResponse, error) {
	id := req.Id

	availableDateDto := mapToAvailableDateDTO(req.AvailableDatedto)

	err := handler.AvailableDateService.Update(id, availableDateDto)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &pb.UpdateAvailableDateResponse{}, nil
}

func (handler *AccommodationHandler) TimeSlotAvailableForAccommodation(ctx context.Context, req *pb.TimeSlotAvailableRequest) (*pb.TimeSlotAvailableResponse, error) {
	dto := mapToTimeSlotAvailableDTO(req.AvailableTimeSlotDTO)

	timeSlotAvailable, err := handler.AvailableDateService.TimeSlotAvailableForAccommodation(dto)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	response := &pb.TimeSlotAvailableResponse{}
	response.Available = timeSlotAvailable
	return response, nil
}

func (handler *AccommodationHandler) GetAutomaticAcceptById(ctx context.Context, req *pb.GetByIdRequest) (*pb.GetAutomaticAcceptByIdResponse, error) {
	id := req.Id
	log.Printf("AvailableDate with id %s", id)
	automaticAccept, err := handler.AccommodationService.GetAutomaticAcceptById(id)

	if err != nil {
		return nil, err
	}
	response := &pb.GetAutomaticAcceptByIdResponse{AutomaticAccept: automaticAccept}
	return response, nil
}
