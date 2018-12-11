package freshdesk

import (
	"fmt"
	"time"
)

type SolutionManager interface {
	Categories() (CategorySlice, error)
}

type solutionManager struct {
	client *apiClient
}

func newSolutionManager(client *apiClient) solutionManager {
	return solutionManager{
		client,
	}
}

type Category struct {
	ID               int       `json:"id,omitempty"`
	Name             string    `json:"name,omitempty"`
	Description      string    `json:"description,omitempty"`
	VisibleInPortals []int     `json:"visible_in_portals"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	client           *apiClient
}

type CategorySlice []Category

func (slice CategorySlice) Len() int {
	return len(slice)
}

func (slice CategorySlice) Less(i, j int) bool {
	return slice[i].ID < slice[j].ID
}

func (slice CategorySlice) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func (slice CategorySlice) Print() {
	for _, category := range slice {
		fmt.Println(category.Name)
	}
}

type Folder struct {
	ID          int       `json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	Description string    `json:"description,omitempty"`
	Visibility  int       `json:"visibility"`
	CompanyIDs  []int     `json:"company_ids"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	client      *apiClient
}

type FolderSlice []Folder

func (slice FolderSlice) Len() int {
	return len(slice)
}

func (slice FolderSlice) Less(i, j int) bool {
	return slice[i].ID < slice[j].ID
}

func (slice FolderSlice) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func (slice FolderSlice) Print() {
	for _, folder := range slice {
		fmt.Println(folder.Name)
	}
}

type Article struct {
	ID              int               `json:"id,omitempty"`
	Title           string            `json:"title"`
	Description     string            `json:"description,omitempty"`
	DescriptionText string            `json:"description_text"`
	AgentID         int               `json:"agent_id"`
	CategoryID      int               `json:"updated_at"`
	FolderID        int               `json:"folder_id"`
	Hits            int               `json:"hits"`
	Status          int               `json:"status"`
	SEOData         map[string]string `json:"seo_data"`
	Tags            []string          `json:"tags"`
	ThumbsDown      int               `json:"thumbs_down"`
	ThumbsUp        int               `json:"thumbs_up"`
	Type            int               `json:"updated_at"`
	CreatedAt       time.Time         `json:"created_at"`
	UpdatedAt       time.Time         `json:"updated_at"`
	client          *apiClient
}

type ArticleSlice []Article

func (slice ArticleSlice) Len() int {
	return len(slice)
}

func (slice ArticleSlice) Less(i, j int) bool {
	return slice[i].ID < slice[j].ID
}

func (slice ArticleSlice) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func (slice ArticleSlice) Print() {
	for _, article := range slice {
		fmt.Println(article.Title)
	}
}

func (manager solutionManager) Categories() (CategorySlice, error) {
	output := CategorySlice{}
	_, err := manager.client.get(endpoints.solutions.categories, &output)
	if err != nil {
		return CategorySlice{}, err
	}
	outputWithClient := CategorySlice{}
	for _, category := range output {
		category.client = manager.client
		outputWithClient = append(outputWithClient, category)
	}
	return outputWithClient, nil
}

func (category Category) Folders() (FolderSlice, error) {
	output := FolderSlice{}
	_, err := category.client.get(endpoints.solutions.category.folders(category.ID), &output)
	if err != nil {
		return FolderSlice{}, err
	}
	outputWithClient := FolderSlice{}
	for _, folder := range output {
		folder.client = category.client
		outputWithClient = append(outputWithClient, folder)
	}
	return outputWithClient, nil
}

func (folder Folder) Articles() (ArticleSlice, error) {
	output := ArticleSlice{}
	_, err := folder.client.get(endpoints.solutions.folder.articles(folder.ID), &output)
	if err != nil {
		return ArticleSlice{}, err
	}
	outputWithClient := ArticleSlice{}
	for _, article := range output {
		article.client = folder.client
		outputWithClient = append(outputWithClient, article)
	}
	return outputWithClient, nil
}

func (article Article) Delete() error {
	return article.client.delete(endpoints.solutions.articles.delete(article.ID))
}
