package controllers

import (
	"context"
	"fmt"
	"TaipeiCityDashboardBE/logs"
	"TaipeiCityDashboardBE/app/models"
	"encoding/json"

	"github.com/mark3labs/mcp-go/mcp"
    "github.com/mark3labs/mcp-go/server"
)

func Activate() {
	s := server.NewMCPServer(
        "TaipeiCityDashboard",
        "1.0.0",
        server.WithResourceCapabilities(true, true),
        server.WithLogging(),
        server.WithRecovery(),
		server.WithInstructions("Get an overview of the dashboard before any moves.\nTables: # components 儀表板組件 儲存每個 dashboard 上的組件（如圖表、地圖等），包含 id、index、name。# component_charts 組件圖表設定 定義每個組件的圖表顏色、型態（如長條圖、圓餅圖）、單位等。# component_maps 組件地圖設定 定義地圖類組件的圖層、樣式、屬性等。# query_charts 儀表板查詢設定 儲存每個圖表的 SQL 查詢語句、說明、資料來源、適用城市、更新頻率等，讓前端能動態產生圖表。"),
    )

overviewTool := mcp.NewTool("overview", mcp.WithDescription("Get an overview of the dashboard"))
s.AddTool(overviewTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// Use the encapsulated service to get all query charts
	db := models.DBManager
	if db == nil {
		return mcp.NewToolResultError("DBManager connection is nil"), nil
	}
	sqldb, err := db.DB()
	if err != nil {
		return mcp.NewToolResultError("Failed to get sql.DB: " + err.Error()), nil
	}
	charts, err := models.GetAllQueryCharts(sqldb)
	if err != nil {
		return mcp.NewToolResultError("GetAllQueryCharts error: " + err.Error()), nil
	}
	jsonBytes, err := json.Marshal(charts)
	if err != nil {
		return mcp.NewToolResultError("JSON marshal error: " + err.Error()), nil
	}
	return mcp.NewToolResultText(string(jsonBytes)), nil
})
		





    // Add a calculator tool
    // calculatorTool := mcp.NewTool("calculate",
    //     mcp.WithDescription("Perform basic arithmetic operations"),
    //     mcp.WithString("operation",
    //         mcp.Required(),
    //         mcp.Description("The operation to perform (add, subtract, multiply, divide)"),
    //         mcp.Enum("add", "subtract", "multiply", "divide"),
    //     ),
    //     mcp.WithNumber("x",
    //         mcp.Required(),
    //         mcp.Description("First number"),
    //     ),
    //     mcp.WithNumber("y",
    //         mcp.Required(),
    //         mcp.Description("Second number"),
    //     ),
    // )

    // // Add the calculator handler
    // s.AddTool(calculatorTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
    //     op := request.Params.Arguments["operation"].(string)
    //     x := request.Params.Arguments["x"].(float64)
    //     y := request.Params.Arguments["y"].(float64)

    //     var result float64
    //     switch op {
    //     case "add":
    //         result = x + y
    //     case "subtract":
    //         result = x - y
    //     case "multiply":
    //         result = x * y
    //     case "divide":
    //         if y == 0 {
    //             return mcp.NewToolResultError("cannot divide by zero"), nil
    //         }
    //         result = x / y
    //     }

    //     return mcp.NewToolResultText(fmt.Sprintf("%.2f", result)), nil
    // })

    // Start the server
	sse := server.NewSSEServer(s, server.WithBaseURL("http://localhost:8081"))
	logs.Info("SSE server listening on :8081")
	if err := sse.Start(":8081"); err != nil {
		panic(fmt.Sprintf("Server error: %v", err))
	}
}
