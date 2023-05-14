package handler

import (
	"Accomodation-Service/domain"
	"Accomodation-Service/dto"
	pb "github.com/XWS-tim24/Common/common/proto/accommodation_service"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func mapToPbAccommodation(accommodation *domain.Accommodation) *pb.Accommodation {
	accPb := &pb.Accommodation{
		Id:        accommodation.Id.String(),
		UserId:    accommodation.UserID,
		Name:      accommodation.Name,
		City:      accommodation.City,
		Address:   accommodation.Address,
		MinGuests: accommodation.MinGuests,
		MaxGuests: accommodation.MaxGuests,
	}

	return accPb
}

func mapToAccommodation(accPb *pb.Accommodation) *domain.Accommodation {
	acc := &domain.Accommodation{
		UserID:    accPb.UserId,
		Name:      accPb.Name,
		City:      accPb.City,
		Address:   accPb.Address,
		MinGuests: accPb.MinGuests,
		MaxGuests: accPb.MaxGuests,
	}

	return acc
}

func mapToAvailableDate(availablePb *pb.AvailableDate) *domain.AvailableDate {
	acc := &domain.AvailableDate{
		AccommodationId: availablePb.Accommodation,
		StartDate:       availablePb.StartDate.AsTime(),
		EndDate:         availablePb.EndDate.AsTime(),
		Price:           uint16(availablePb.Price),
		PricingType:     mapToPricingType(availablePb.PricingType),
	}

	return acc
}

func mapToAvailableDatePb(available *domain.AvailableDate) *pb.AvailableDate {
	availablePb := &pb.AvailableDate{
		Id:            available.Id.String(),
		Accommodation: available.AccommodationId,
		StartDate:     timestamppb.New(available.StartDate),
		EndDate:       timestamppb.New(available.EndDate),
		Price:         uint32(available.Price),
		PricingType:   mapToPricingTypePb(available.PricingType),
	}

	return availablePb
}

func mapToPricingType(pricingType pb.AvailableDate_PricingType) domain.PricingType {
	if pricingType == pb.AvailableDate_Per_Guest {
		return domain.PER_GUEST
	}
	return domain.PER_ACCOMMODATION
}

func mapToPricingTypePb(pricingType domain.PricingType) pb.AvailableDate_PricingType {
	if pricingType == domain.PER_GUEST {
		return pb.AvailableDate_Per_Guest
	}
	return pb.AvailableDate_Per_Accommodation
}

func mapAccommodationSearchDTO(searchDto *pb.AccommodationSearchDTO) *dto.AccommodationSearchDTO {
	acc := &dto.AccommodationSearchDTO{
		City:      searchDto.City,
		GuestNum:  uint16(searchDto.GuestNum),
		StartDate: searchDto.StartDate.AsTime(),
		EndDate:   searchDto.EndDate.AsTime(),
	}

	return acc
}

func mapToAvailableDateDTO(pbUpdateDTO *pb.AvailableDateDTO) *dto.AvailableDateDTO {

	updateDto := &dto.AvailableDateDTO{
		Id:          pbUpdateDTO.Id,
		StartDate:   pbUpdateDTO.StartDate.AsTime(),
		EndDate:     pbUpdateDTO.EndDate.AsTime(),
		Price:       uint16(pbUpdateDTO.Price),
		PricingType: mapToPricingTypeDTO(pbUpdateDTO.PricingType),
	}

	return updateDto
}

func mapToPricingTypeDTO(pricingType pb.AvailableDateDTO_PricingType) domain.PricingType {
	if pricingType == pb.AvailableDateDTO_Per_Guest {
		return domain.PER_GUEST
	}
	return domain.PER_ACCOMMODATION
}
