//
// Copyright 2023, Nick Westbury
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package gitlab

import (
	"fmt"
	"net/http"
	"time"
)

// ProjectRepositoryStorageMoveService handles communication with the
// repositories related methods of the GitLab API.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/project_repository_storage_moves.html
type ProjectRepositoryStorageMoveService struct {
	client *Client
}

// ProjectRepositoryStorageMove represents the status of a repository move.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/project_repository_storage_moves.html
type ProjectRepositoryStorageMove struct {
	ID                     int        `json:"id"`
	CreatedAt              *time.Time `json:"created_at"`
	State                  string     `json:"state"`
	SourceStorageName      string     `json:"source_storage_name"`
	DestinationStorageName string     `json:"destination_storage_name"`
	Project                struct {
		ID                int        `json:"id"`
		Description       string     `json:"description"`
		Name              string     `json:"name"`
		NameWithNamespace string     `json:"name_with_namespace"`
		Path              string     `json:"path"`
		PathWithNamespace string     `json:"path_with_namespace"`
		CreatedAt         *time.Time `json:"created_at"`
	} `json:"project"`
}

// RetrieveAllStorageMovesOptions represents the available
// RetrieveAllStorageMoves() options.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/project_repository_storage_moves.html
type RetrieveAllStorageMovesOptions ListOptions

// RetrieveAllStorageMoves retrieves all repository storage moves accessible by
// the authenticated user.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/project_repository_storage_moves.html#retrieve-all-project-repository-storage-moves
func (s ProjectRepositoryStorageMoveService) RetrieveAllStorageMoves(opts RetrieveAllStorageMovesOptions, options ...RequestOptionFunc) ([]*ProjectRepositoryStorageMove, *Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "project_repository_storage_moves", opts, options)
	if err != nil {
		return nil, nil, err
	}

	var psms []*ProjectRepositoryStorageMove
	resp, err := s.client.Do(req, &psms)
	if err != nil {
		return nil, resp, err
	}

	return psms, resp, err
}

// RetrieveAllStorageMovesForProject retrieves all repository storage moves for
// a single project accessible by the authenticated user.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/project_repository_storage_moves.html#retrieve-all-repository-storage-moves-for-a-project
func (s ProjectRepositoryStorageMoveService) RetrieveAllStorageMovesForProject(project int, opts RetrieveAllStorageMovesOptions, options ...RequestOptionFunc) ([]*ProjectRepositoryStorageMove, *Response, error) {
	u := fmt.Sprintf("projects/%d/repository_storage_moves", project)

	req, err := s.client.NewRequest(http.MethodGet, u, opts, options)
	if err != nil {
		return nil, nil, err
	}

	var psms []*ProjectRepositoryStorageMove
	resp, err := s.client.Do(req, &psms)
	if err != nil {
		return nil, resp, err
	}

	return psms, resp, err
}

// GetStorageMove gets a single repository storage move.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/project_repository_storage_moves.html#get-a-single-project-repository-storage-move
func (s ProjectRepositoryStorageMoveService) GetStorageMove(repositoryStorage int, options ...RequestOptionFunc) (*ProjectRepositoryStorageMove, *Response, error) {
	u := fmt.Sprintf("project_repository_storage_moves/%d", repositoryStorage)

	req, err := s.client.NewRequest(http.MethodGet, u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	psm := new(ProjectRepositoryStorageMove)
	resp, err := s.client.Do(req, psm)
	if err != nil {
		return nil, resp, err
	}

	return psm, resp, err
}

// GetStorageMoveForProject gets a single repository storage move for a project.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/project_repository_storage_moves.html#get-a-single-repository-storage-move-for-a-project
func (s ProjectRepositoryStorageMoveService) GetStorageMoveForProject(project int, repositoryStorage int, options ...RequestOptionFunc) (*ProjectRepositoryStorageMove, *Response, error) {
	u := fmt.Sprintf("projects/%d/repository_storage_moves/%d", project, repositoryStorage)

	req, err := s.client.NewRequest(http.MethodGet, u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	psm := new(ProjectRepositoryStorageMove)
	resp, err := s.client.Do(req, psm)
	if err != nil {
		return nil, resp, err
	}

	return psm, resp, err
}

// ScheduleAllStorageMoves schedules all repositories to be moved.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/project_repository_storage_moves.html#schedule-repository-storage-moves-for-all-projects-on-a-storage-shard
func (s ProjectRepositoryStorageMoveService) ScheduleAllStorageMoves(options ...RequestOptionFunc) ([]*ProjectRepositoryStorageMove, *Response, error) {
	req, err := s.client.NewRequest(http.MethodPost, "project_repository_storage_moves", nil, options)
	if err != nil {
		return nil, nil, err
	}

	var psms []*ProjectRepositoryStorageMove
	resp, err := s.client.Do(req, &psms)
	if err != nil {
		return nil, resp, err
	}

	return psms, resp, err
}

// ScheduleStorageMoveForProject schedule a repository to be moved for a project.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/project_repository_storage_moves.html#schedule-a-repository-storage-move-for-a-project
func (s ProjectRepositoryStorageMoveService) ScheduleStorageMoveForProject(project int, options ...RequestOptionFunc) ([]*ProjectRepositoryStorageMove, *Response, error) {
	u := fmt.Sprintf("projects/%d/repository_storage_moves", project)

	req, err := s.client.NewRequest(http.MethodPost, u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	var psms []*ProjectRepositoryStorageMove
	resp, err := s.client.Do(req, &psms)
	if err != nil {
		return nil, resp, err
	}

	return psms, resp, err
}
