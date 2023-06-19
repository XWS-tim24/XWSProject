package handler

import (
	"Rating-Service/domain"
	"time"

	pb "github.com/XWS-tim24/Common/common/proto/rating_service"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func mapToHostRatingPb(rating *domain.HostRating) *pb.HostRating {
	ratingPb := &pb.HostRating{
		Id:     rating.Id.String(),
		UserId: rating.UserId,
		HostId: rating.HostId,
		Date:   timestamppb.New(rating.Date),
		Value:  rating.Value,
	}
	return ratingPb
}

func mapToNewHostRating(ratingPb *pb.NewHostRating) *domain.HostRating {
	rating := &domain.HostRating{
		HostId: ratingPb.HostId,
		UserId: ratingPb.UserId,
		Value:  ratingPb.Value,
		Date:   time.Now(),
	}
	return rating
}

func mapToAccommodationRatingPb(rating *domain.AccommodationRating) *pb.AccommodationRating {
	ratingPb := &pb.AccommodationRating{
		Id:              rating.Id.String(),
		UserId:          rating.UserId,
		AccommodationId: rating.AccommodationId,
		Date:            timestamppb.New(rating.Date),
		Value:           rating.Value,
	}
	return ratingPb
}

func mapToNewAccommodationRating(ratingPb *pb.NewAccommodationRating) *domain.AccommodationRating {
	rating := &domain.AccommodationRating{
		AccommodationId: ratingPb.AccommodationId,
		UserId:          ratingPb.UserId,
		Value:           ratingPb.Value,
		Date:            time.Now(),
	}
	return rating
}

/*
func mapToNewHostRatingPb(rating *domain.HostRating) *pb.NewHostRating {
	ratingPb := &pb.NewHostRating{
		HostId: rating.HostId,
		UserId: rating.UserId,
		Value:  rating.Value,
	}
	return ratingPb
}*/

/*
func mapToHostRating(ratingPb *pb.HostRating) *domain.HostRating {
	rating := &domain.HostRating{
		HostId: ratingPb.HostId,
		UserId: ratingPb.UserId,
	}
	return rating
}*/
