package mapper

import (
	"Accomodation-Service/domain"
	"Accomodation-Service/dto"

	pb "github.com/XWS-tim24/Common/common/proto/accommodation_service"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func MapToPbAccommodation(accommodation *domain.Accommodation) *pb.Accommodation {
	accPb := &pb.Accommodation{
		Id:              accommodation.Id.String(),
		UserId:          accommodation.UserID,
		Name:            accommodation.Name,
		City:            accommodation.City,
		Address:         accommodation.Address,
		Benefits:        accommodation.Benefits,
		MinGuests:       accommodation.MinGuests,
		MaxGuests:       accommodation.MaxGuests,
		AutomaticAccept: accommodation.AutomaticAccept,
	}

	return accPb
}

func MapToAccommodation(accPb *pb.Accommodation) *domain.Accommodation {
	acc := &domain.Accommodation{
		UserID:          accPb.UserId,
		Name:            accPb.Name,
		City:            accPb.City,
		Address:         accPb.Address,
		Benefits:        accPb.Benefits,
		MinGuests:       accPb.MinGuests,
		MaxGuests:       accPb.MaxGuests,
		AutomaticAccept: accPb.AutomaticAccept,
	}

	return acc
}

func MapToAvailableDate(availablePb *pb.AvailableDate) *domain.AvailableDate {
	acc := &domain.AvailableDate{
		AccommodationId: availablePb.Accommodation,
		StartDate:       availablePb.StartDate.AsTime(),
		EndDate:         availablePb.EndDate.AsTime(),
		Price:           uint16(availablePb.Price),
		PricingType:     MapToPricingType(availablePb.PricingType),
	}

	return acc
}

func MapToAvailableDatePb(available *domain.AvailableDate) *pb.AvailableDate {
	availablePb := &pb.AvailableDate{
		Id:            available.Id.String(),
		Accommodation: available.AccommodationId,
		StartDate:     timestamppb.New(available.StartDate),
		EndDate:       timestamppb.New(available.EndDate),
		Price:         uint32(available.Price),
		PricingType:   MapToPricingTypePb(available.PricingType),
	}

	return availablePb
}

func MapToPricingType(pricingType pb.AvailableDate_PricingType) domain.PricingType {
	if pricingType == pb.AvailableDate_Per_Guest {
		return domain.PER_GUEST
	}
	return domain.PER_ACCOMMODATION
}

func MapToPricingTypePb(pricingType domain.PricingType) pb.AvailableDate_PricingType {
	if pricingType == domain.PER_GUEST {
		return pb.AvailableDate_Per_Guest
	}
	return pb.AvailableDate_Per_Accommodation
}

func MapAccommodationSearchDTO(searchDto *pb.AccommodationSearchDTO) *dto.AccommodationSearchDTO {
	acc := &dto.AccommodationSearchDTO{
		City:      searchDto.City,
		GuestNum:  uint16(searchDto.GuestNum),
		StartDate: searchDto.StartDate.AsTime(),
		EndDate:   searchDto.EndDate.AsTime(),
	}

	return acc
}

func MapToAvailableDateDTO(pbUpdateDTO *pb.AvailableDateDTO) *dto.AvailableDateDTO {

	updateDto := &dto.AvailableDateDTO{
		Id:          pbUpdateDTO.Id,
		StartDate:   pbUpdateDTO.StartDate.AsTime(),
		EndDate:     pbUpdateDTO.EndDate.AsTime(),
		Price:       uint16(pbUpdateDTO.Price),
		PricingType: MapToPricingTypeDTO(pbUpdateDTO.PricingType),
	}

	return updateDto
}

func MapToPricingTypeDTO(pricingType pb.AvailableDateDTO_PricingType) domain.PricingType {
	if pricingType == pb.AvailableDateDTO_Per_Guest {
		return domain.PER_GUEST
	}
	return domain.PER_ACCOMMODATION
}

func MapToTimeSlotAvailableDTO(pbUTimeSlotDTO *pb.AvailableTimeSlotDTO) *dto.AvailableTimeSlotDTO {

	updateDto := &dto.AvailableTimeSlotDTO{
		AccommodationId: pbUTimeSlotDTO.AccommodationId,
		StartDate:       pbUTimeSlotDTO.StartDate.AsTime(),
		EndDate:         pbUTimeSlotDTO.EndDate.AsTime(),
	}

	return updateDto
}

func MapToSearchResponse(accommodation *domain.Accommodation, price uint16, totalPrice uint16, pricingType domain.PricingType) *dto.AccommodationSearchDTOResponse {
	accSearchResponse := &dto.AccommodationSearchDTOResponse{
		Id:          accommodation.Id.String(),
		Name:        accommodation.Name,
		City:        accommodation.City,
		Address:     accommodation.Address,
		Benefits:    accommodation.Benefits,
		MinGuests:   accommodation.MinGuests,
		MaxGuests:   accommodation.MaxGuests,
		Price:       price,
		TotalPrice:  totalPrice,
		PricingType: pricingType,
	}
	return accSearchResponse
}

func MapSearchResponseToPbResponse(searchResponse *dto.AccommodationSearchDTOResponse) *pb.AccommodationDTOForSearchResponse {
	retVal := &pb.AccommodationDTOForSearchResponse{
		Id:          searchResponse.Id,
		Name:        searchResponse.Name,
		City:        searchResponse.City,
		Address:     searchResponse.Address,
		Benefits:    searchResponse.Benefits,
		MinGuests:   searchResponse.MinGuests,
		MaxGuests:   searchResponse.MaxGuests,
		Price:       uint32(searchResponse.Price),
		TotalPrice:  uint32(searchResponse.TotalPrice),
		PricingType: MapPricingType(searchResponse.PricingType),
	}

	return retVal
}

func MapPricingType(pricingType domain.PricingType) pb.AccommodationDTOForSearchResponse_PricingType {
	if pricingType == domain.PER_GUEST {
		return pb.AccommodationDTOForSearchResponse_Per_Guest
	}
	return pb.AccommodationDTOForSearchResponse_Per_Accommodation
}
