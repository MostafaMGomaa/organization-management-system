package memberRepositroy

import (
	"context"
	"ideanest/pkg/database/mongodb/models"
)

type MemberRepository interface {
	Createmember(ctx context.Context, member models.Members) error
}

type MongoDBmemberRepository struct {
	collection *mongo.Collection
}

// constructor function
func NewMongoDBMemberRepository(collection *mongo.Collection) *MongoDBmemberRepository {
	return &MongoDBmemberRepository{
		collection: collection,
	}
}

func (r *MongoDBmemberRepository) CreateMember(ctx context.Context, member models.Members) error {
	memberDocument, err := r.collection.InsertOne(ctx, member)
	if err != nil {
		return err
	}
	return memberDocument
}
