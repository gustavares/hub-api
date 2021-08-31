package usecase

import (
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gustavares/hub-api/internal/domain/entity"
	"github.com/gustavares/hub-api/internal/repository"
)

func TestGetUserPosts_Do(t *testing.T) {
	userId := "123"

	post1 := entity.NewPost("1", "video", "test")
	post2 := entity.NewPost("2", "photo", "test")
	posts := []entity.Post{
		*post1,
		*post2,
	}

	repoMock := repository.NewMockPostRepository(gomock.NewController(t))
	repoMock.EXPECT().GetByUser(userId).Times(1).Return(&posts, nil)

	emptyRepoMock := repository.NewMockPostRepository(gomock.NewController(t))
	emptyRepoMock.EXPECT().GetByUser("555").Times(1).Return(&[]entity.Post{}, nil)

	type fields struct {
		postRepository repository.PostRepository
	}
	type args struct {
		userId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *[]entity.Post
		wantErr bool
	}{
		{
			name: "given a user ID it should return all of its posts and no errors",
			fields: fields{
				postRepository: repoMock,
			},
			args: args{
				userId: userId,
			},
			want:    &posts,
			wantErr: false,
		},
		{
			name: "given a user ID that does not have any posts it should return an error",
			fields: fields{
				postRepository: emptyRepoMock,
			},
			args: args{
				userId: "555",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GetUserPosts{
				postRepository: tt.fields.postRepository,
			}
			got, err := g.Do(tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserPosts.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserPosts.Do() = %v, want %v", got, tt.want)
			}
		})
	}
}
