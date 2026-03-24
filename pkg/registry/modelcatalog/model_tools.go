package modelcatalog

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/digitalocean/godo"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// ModelTool provides model catalog management tools
type ModelTool struct {
	client func(ctx context.Context) (*godo.Client, error)
}

// NewModelTool creates a new ModelTool instance
func NewModelTool(client func(ctx context.Context) (*godo.Client, error)) *ModelTool {
	return &ModelTool{
		client: client,
	}
}

// searchModels searches for models in the catalog using a search string
func (m *ModelTool) searchModels(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	searchQuery, ok := req.GetArguments()["SearchQuery"].(string)
	if !ok || searchQuery == "" {
		return mcp.NewToolResultError("SearchQuery is required"), nil
	}

	client, err := m.client(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get DigitalOcean client: %w", err)
	}

	// TODO: Call GoDo API for catalog search
	// Example: models, _, err := client.ModelCatalog.Search(ctx, searchQuery)
	_ = client

	// Placeholder response structure
	type ModelSearchResult struct {
		ModelIDs []string `json:"model_ids"`
		Query    string   `json:"query"`
	}

	result := ModelSearchResult{
		ModelIDs: []string{},
		Query:    searchQuery,
	}

	jsonData, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("marshal error: %w", err)
	}

	return mcp.NewToolResultText(string(jsonData)), nil
}

// getModelCard retrieves full metadata for a specific model
func (m *ModelTool) getModelCard(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	modelID, ok := req.GetArguments()["ModelID"].(string)
	if !ok || modelID == "" {
		return mcp.NewToolResultError("ModelID is required"), nil
	}

	client, err := m.client(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get DigitalOcean client: %w", err)
	}

	// TODO: Call GoDo API to get model card
	// Example: modelCard, _, err := client.ModelCatalog.Get(ctx, modelID)
	_ = client

	// Placeholder response structure
	type ModelCard struct {
		ModelID        string   `json:"model_id"`
		Name           string   `json:"name"`
		Description    string   `json:"description"`
		LicenseType    string   `json:"license_type"`
		AttributionURL string   `json:"attribution_url"`
		DeploymentType []string `json:"deployment_type"`
	}

	modelCard := ModelCard{
		ModelID:        modelID,
		Name:           "",
		Description:    "",
		LicenseType:    "",
		AttributionURL: "",
		DeploymentType: []string{},
	}

	jsonData, err := json.MarshalIndent(modelCard, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("marshal error: %w", err)
	}

	return mcp.NewToolResultText(string(jsonData)), nil
}

// Tools returns the list of server tools for model catalog management
func (m *ModelTool) Tools() []server.ServerTool {
	return []server.ServerTool{
		{
			Handler: m.searchModels,
			Tool: mcp.NewTool(
				"model-catalog-search",
				mcp.WithDescription("Search for models in the catalog using a search query. Returns a list of model IDs that match the search criteria."),
				mcp.WithString("SearchQuery", mcp.Required(), mcp.Description("Search query string to find models")),
			),
		},
		{
			Handler: m.getModelCard,
			Tool: mcp.NewTool(
				"model-catalog-get-card",
				mcp.WithDescription("Get the model metadata for a specific model ID."),
				mcp.WithString("ModelID", mcp.Required(), mcp.Description("The unique identifier of the model")),
			),
		},
	}
}
