package handler

import (
	"Accomodation-Service/mapper"
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
	accommodation := mapper.MapToAccommodation(pbAcc.Accommodation)
	println("Body:")

	acc, err := handler.AccommodationService.Create(accommodation)
	if err != nil {
		return nil, err
	}
	resp := &pb.CreateAccommodationResponse{}
	resp.Accommodation = mapper.MapToPbAccommodation(acc)
	return resp, nil
}

func (handler *AccommodationHandler) GetAccommodationById(ctx context.Context, req *pb.GetByIdRequest) (*pb.CreateAccommodationResponse, error) {
	id := req.Id

	println("Accommodation with id %s", id)
	accommodation, err := handler.AccommodationService.GetById(id)

	if err != nil {
		return nil, err
	}
	pbAccommodation := mapper.MapToPbAccommodation(accommodation)
	return &pb.CreateAccommodationResponse{Accommodation: pbAccommodation}, nil
}

func (handler *AccommodationHandler) SearchAccommodations(ctx context.Context, req *pb.SearchRequest) (*pb.SearchResponse, error) {

	accommodationSearchDTO := mapper.MapAccommodationSearchDTO(req.AccommodationSearchDTO)
	println("Body:")
	accStr, _ := json.Marshal(accommodationSearchDTO)
	println(string(accStr))

	searchResponses, err := handler.AccommodationService.Search(accommodationSearchDTO)
	if err != nil {
		return nil, err
	}
	response := &pb.SearchResponse{
		AccommodationDto: []*pb.AccommodationDTOForSearchResponse{},
	}
	for _, acc := range *searchResponses {
		current := mapper.MapSearchResponseToPbResponse(&acc)
		response.AccommodationDto = append(response.AccommodationDto, current)
	}
	return response, nil
}

func (handler *AccommodationHandler) CreateAvailableDate(ctx context.Context, pbAvailableDate *pb.CreateAvailableDateRequest) (*pb.CreateAvailableDateResponse, error) {

	availableDate := mapper.MapToAvailableDate(pbAvailableDate.AvailableDate)

	availableDate, err := handler.AvailableDateService.Create(availableDate)
	if err != nil {
		println("Error while creating a new Available Date")
		return nil, err
	}
	resp := pb.CreateAvailableDateResponse{}
	resp.AvailableDate = mapper.MapToAvailableDatePb(availableDate)
	return &resp, nil
}

func (handler *AccommodationHandler) GetAvailableDateById(ctx context.Context, req *pb.GetByIdRequest) (*pb.CreateAvailableDateResponse, error) {
	id := req.Id
	log.Printf("AvailableDate with id %s", id)
	availableDate, err := handler.AvailableDateService.GetById(id)

	if err != nil {
		return nil, err
	}

	availableDatePb := mapper.MapToAvailableDatePb(availableDate)
	return &pb.CreateAvailableDateResponse{AvailableDate: availableDatePb}, nil
}

func (handler *AccommodationHandler) UpdateAvailableDate(ctx context.Context, req *pb.UpdateAvailableDateRequest) (*pb.UpdateAvailableDateResponse, error) {
	id := req.Id

	availableDateDto := mapper.MapToAvailableDateDTO(req.AvailableDatedto)

	err := handler.AvailableDateService.Update(id, availableDateDto)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &pb.UpdateAvailableDateResponse{}, nil
}

func (handler *AccommodationHandler) TimeSlotAvailableForAccommodation(ctx context.Context, req *pb.TimeSlotAvailableRequest) (*pb.TimeSlotAvailableResponse, error) {
	dto := mapper.MapToTimeSlotAvailableDTO(req.AvailableTimeSlotDTO)
	fmt.Println("Is Time slot availble :id", dto.AccommodationId)
	fmt.Println(" startDate:", dto.StartDate)
	fmt.Println(" Enddate:", dto.EndDate)

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

func (handler *AccommodationHandler) GetAllAccommodation(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllAccommodationResponse, error) {
	accomodations, err := handler.AccommodationService.GetAll()
	if err != nil || *accomodations == nil {
		return nil, err
	}
	response := &pb.GetAllAccommodationResponse{
		Accommodations: []*pb.Accommodation{},
	}
	for _, reservation := range *accomodations {
		current := mapper.MapToPbAccommodation(&reservation)
		response.Accommodations = append(response.Accommodations, current)
	}
	return response, nil
}

func (handler *AccommodationHandler) GetAllAvailableDates(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllAvailableDatesResponse, error) {
	dates, err := handler.AvailableDateService.GetAll()
	if err != nil || *dates == nil {
		return nil, err
	}
	response := &pb.GetAllAvailableDatesResponse{
		AvailableDates: []*pb.AvailableDate{},
	}
	for _, reservationRequest := range *dates {
		current := mapper.MapToAvailableDatePb(&reservationRequest)
		response.AvailableDates = append(response.AvailableDates, current)
	}
	return response, nil
}
