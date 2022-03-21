package bybit

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// CreateLinearOrderResponse :
type CreateLinearOrderResponse struct {
	CommonResponse `json:",inline"`
	Result         CreateLinearOrderResult `json:"result"`
}

// CreateLinearOrderResult :
type CreateLinearOrderResult struct {
	CreateLinearOrder `json:",inline"`
}

// CreateLinearOrder :
type CreateLinearOrder struct {
	OrderID        string      `json:"order_id"`
	UserID         int         `json:"user_id"`
	Symbol         SymbolUSDT  `json:"symbol"`
	Side           Side        `json:"side"`
	OrderType      OrderType   `json:"order_type"`
	Price          float64     `json:"price"`
	Qty            float64     `json:"qty"`
	TimeInForce    TimeInForce `json:"time_in_force"`
	OrderStatus    OrderStatus `json:"order_status"`
	LastExecPrice  float64     `json:"last_exec_price"`
	CumExecQty     float64     `json:"cum_exec_qty"`
	CumExecValue   float64     `json:"cum_exec_value"`
	CumExecFee     float64     `json:"cum_exec_fee"`
	ReduceOnly     bool        `json:"reduce_only"`
	CloseOnTrigger bool        `json:"close_on_trigger"`
	OrderLinkID    string      `json:"order_link_id"`
	CreatedTime    string      `json:"created_time"`
	UpdatedTime    string      `json:"updated_time"`
	TakeProfit     float64     `json:"take_profit"`
	StopLoss       float64     `json:"stop_loss"`
	TpTriggerBy    string      `json:"tp_trigger_by"`
	SlTriggerBy    string      `json:"sl_trigger_by"`
}

// CreateLinearOrderParam :
type CreateLinearOrderParam struct {
	Side           Side        `json:"side"`
	Symbol         SymbolUSDT  `json:"symbol"`
	OrderType      OrderType   `json:"order_type"`
	Qty            float64     `json:"qty"`
	TimeInForce    TimeInForce `json:"time_in_force"`
	ReduceOnly     bool        `json:"reduce_only"`
	CloseOnTrigger bool        `json:"close_on_trigger"`

	Price       *float64 `json:"price,omitempty"`
	TakeProfit  *float64 `json:"take_profit,omitempty"`
	StopLoss    *float64 `json:"stop_loss,omitempty"`
	TpTriggerBy *string  `json:"tp_trigger_by"`
	SlTriggerBy *string  `json:"sl_trigger_by"`
	OrderLinkID *string  `json:"order_link_id,omitempty"`
}

// CreateLinearOrder :
func (s *AccountService) CreateLinearOrder(param CreateLinearOrderParam) (*CreateLinearOrderResponse, error) {
	var res CreateLinearOrderResponse

	url, err := s.Client.BuildPrivateURL("/private/linear/order/create", nil)
	if err != nil {
		return nil, err
	}

	jsonBody, err := json.Marshal(param)
	if err != nil {
		return nil, fmt.Errorf("json marshal for CreateLinearOrderParam: %w", err)
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}
	return &res, nil
}

// ListLinearPositionResponse :
type ListLinearPositionResponse struct {
	CommonResponse `json:",inline"`
	Result         []ListLinearPositionResult `json:"result"`
}

// ListLinearPositionResult :
type ListLinearPositionResult struct {
	UserID              int           `json:"user_id"`
	Symbol              SymbolInverse `json:"symbol"`
	Side                Side          `json:"side"`
	Size                float64       `json:"size"`
	PositionValue       float64       `json:"position_value"`
	EntryPrice          float64       `json:"entry_price"`
	LiqPrice            float64       `json:"liq_price"`
	BustPrice           float64       `json:"bust_price"`
	Leverage            float64       `json:"leverage"`
	AutoAddMargin       float64       `json:"auto_add_margin"`
	IsIsolated          bool          `json:"is_isolated"`
	PositionMargin      float64       `json:"position_margin"`
	OccClosingFee       float64       `json:"occ_closing_fee"`
	RealisedPnl         float64       `json:"realised_pnl"`
	CumRealisedPnl      float64       `json:"cum_realised_pnl"`
	FreeQty             float64       `json:"free_qty"`
	TpSlMode            TpSlMode      `json:"tp_sl_mode"`
	DeleverageIndicator int           `json:"deleverage_indicator"`
	UnrealisedPnl       float64       `json:"unrealised_pnl"`
	RiskID              int           `json:"risk_id"`
}

// ListLinearPosition :
func (s *AccountService) ListLinearPosition(symbol SymbolUSDT) (*ListLinearPositionResponse, error) {
	var res ListLinearPositionResponse

	params := map[string]string{
		"symbol": string(symbol),
	}
	url, err := s.Client.BuildPrivateURL("/private/linear/position/list", params)
	if err != nil {
		return nil, err
	}
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}
	return &res, nil
}

// ListLinearPositionsResponse :
type ListLinearPositionsResponse struct {
	CommonResponse `json:",inline"`
	Result         []ListLinearPositionsResult `json:"result"`
}

// ListLinearPositionsResult :
type ListLinearPositionsResult struct {
	IsValid                  bool `json:"is_valid"`
	ListLinearPositionResult `json:"data,inline"`
}

// ListLinearPositions :
func (s *AccountService) ListLinearPositions() (*ListLinearPositionsResponse, error) {
	var res ListLinearPositionsResponse

	url, err := s.Client.BuildPrivateURL("/private/linear/position/list", nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}
	return &res, nil
}

// CancelLinearOrderResponse :
type CancelLinearOrderResponse struct {
	CommonResponse `json:",inline"`
	Result         CancelLinearOrderResult `json:"result"`
}

// CancelLinearOrderResult :
type CancelLinearOrderResult struct {
	CancelLinearOrder `json:",inline"`
}

// CancelLinearOrder :
type CancelLinearOrder struct {
	OrderID string `json:"order_id"`
}

// QueryLinearOrderParam :
type QueryLinearOrderParam struct {
	Symbol SymbolUSDT `json:"symbol"`

	OrderID     *string `json:"order_id,omitempty"`
	OrderLinkID *string `json:"order_link_id,omitempty"`
}

// CancelLinearOrder :
func (s *AccountService) CancelLinearOrder(param QueryLinearOrderParam) (*CancelLinearOrderResponse, error) {
	var res CancelLinearOrderResponse

	if param.OrderID == nil && param.OrderLinkID == nil {
		return nil, fmt.Errorf("either OrderID or OrderLinkID needed")
	}

	url, err := s.Client.BuildPrivateURL("/private/linear/order/cancel", nil)
	if err != nil {
		return nil, err
	}

	jsonBody, err := json.Marshal(param)
	if err != nil {
		return nil, fmt.Errorf("json marshal for QueryLinearOrderParam: %w", err)
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}
	return &res, nil
}

// SaveLinearLeverageResponse :
type SaveLinearLeverageResponse struct {
	CommonResponse `json:",inline"`
}

// SaveLinearLeverageParam :
type SaveLinearLeverageParam struct {
	Symbol       SymbolUSDT `json:"symbol"`
	BuyLeverage  float64    `json:"buy_leverage"`
	SellLeverage float64    `json:"sell_leverage"`
}

// SaveLinearLeverage :
func (s *AccountService) SaveLinearLeverage(param SaveLinearLeverageParam) (*SaveLinearLeverageResponse, error) {
	var res SaveLinearLeverageResponse

	url, err := s.Client.BuildPrivateURL("/private/linear/position/set-leverage", nil)
	if err != nil {
		return nil, err
	}

	jsonBody, err := json.Marshal(param)
	if err != nil {
		return nil, fmt.Errorf("json marshal for SaveLinearLeverageParam: %w", err)
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}
	return &res, nil
}

// LinearExecutionListResponse :
type LinearExecutionListResponse struct {
	CommonResponse `json:",inline"`
	Result         LinearExecutionListResult `json:"result"`
}

// LinearExecutionListResult :
type LinearExecutionListResult struct {
	CurrentPage          int                   `json:"current_page"`
	LinearExecutionLists []LinearExecutionList `json:"data"`
}

// LinearExecutionList :
type LinearExecutionList struct {
	OrderID          string     `json:"order_id"`
	OrderLinkID      string     `json:"order_link_id"`
	Side             Side       `json:"side"`
	Symbol           SymbolUSDT `json:"symbol"`
	OrderPrice       float64    `json:"order_price"`
	OrderQty         float64    `json:"order_qty"`
	OrderType        OrderType  `json:"order_type"`
	FeeRate          float64    `json:"fee_rate"`
	ExecPrice        float64    `json:"exec_price"`
	ExecType         ExecType   `json:"exec_type"`
	ExecQty          float64    `json:"exec_qty"`
	ExecFee          float64    `json:"exec_fee"`
	ExecValue        float64    `json:"exec_value"`
	LeavesQty        float64    `json:"leaves_qty"`
	ClosedSize       float64    `json:"closed_size"`
	LastLiquidityInd string     `json:"last_liquidity_ind"`
	TradeTimeMs      float64    `json:"trade_time_ms"`
}

// LinearExecutionListParam :
type LinearExecutionListParam struct {
	Symbol SymbolUSDT `json:"symbol"`

	StartTime *int      `json:"start_time"`
	EndTime   *int      `json:"end_time"`
	ExecType  *ExecType `json:"exec_type"`
	Page      *int      `json:"page"`
	Limit     *int      `json:"limit"`
}

// LinearExecutionList :
// NOTE(TODO) : somehow got EOF 404(path not found)
func (s *AccountService) LinearExecutionList(param LinearExecutionListParam) (*LinearExecutionListResponse, error) {
	var res LinearExecutionListResponse

	url, err := s.Client.BuildPrivateURL("/private/linear/trade/execution/list", nil)
	if err != nil {
		return nil, err
	}

	jsonBody, err := json.Marshal(param)
	if err != nil {
		return nil, fmt.Errorf("json marshal for LinearExecutionListParam: %w", err)
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}
	return &res, nil
}

// LinearOrderListResponse (Get Active Orders) :
type LinearOrderListResponse struct {
	CommonResponse `json:",inline"`
	Result         LinearOrderListResult `json:"result"`
}

// LinearOrderListResult :
type LinearOrderListResult struct {
	CurrentPage     int                 `json:"current_page"`
	LinearOrderList []CreateLinearOrder `json:"data"`
}

// LinearExecutionListParam :
type LinearOrderListParam struct {
	Symbol      SymbolUSDT `json:"symbol"`
	OrderID     *string    `json:"order_id,omitempty"`
	OrderLinkID *string    `json:"order_link_id,omitempty"`
	Order       *string    `json:"order,omitempty"`
	Page        *int       `json:"page,omitempty"`
	Limit       *int       `json:"limit,omitempty"`
	OrderStatus string     `json:"order_status,omitempty"`
}

// LinearOrderList :
func (s *AccountService) LinearOrderList(param LinearOrderListParam) (*LinearOrderListResponse, error) {
	var res LinearOrderListResponse

	url, err := s.Client.BuildPrivateURL("/private/linear/order/list", nil)
	if err != nil {
		return nil, err
	}

	jsonBody, err := json.Marshal(param)
	if err != nil {
		return nil, fmt.Errorf("json marshal for LinearOrderListParam: %w", err)
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}
	return &res, nil
}

// CancelAllLinearOrderResponse :
type CancelAllLinearOrderResponse struct {
	CommonResponse `json:",inline"`
	Result         []string `json:"result"`
}

// CancelAllLinearOrderParam :
type CancelAllLinearOrderParam struct {
	Symbol SymbolUSDT `json:"symbol"`
}

// CancelAllLinearOrder :
func (s *AccountService) CancelAllLinearOrder(param CancelAllLinearOrderParam) (*CancelAllLinearOrderResponse, error) {
	var res CancelAllLinearOrderResponse

	url, err := s.Client.BuildPrivateURL("/private/linear/order/cancel-all", nil)
	if err != nil {
		return nil, err
	}

	jsonBody, err := json.Marshal(param)
	if err != nil {
		return nil, fmt.Errorf("json marshal for CancelAllLinearOrderParam: %w", err)
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}
	return &res, nil
}

// LinearOrderReplaceResponse :
type LinearOrderReplaceResponse struct {
	CancelLinearOrderResponse
}

// LinearOrderReplaceParam :
type LinearOrderReplaceParam struct {
	Symbol                 SymbolUSDT `json:"symbol"`
	OrderID                *string    `json:"order_id,omitempty"`
	OrderLinkID            *string    `json:"order_link_id,omitempty"`
	Amount                 *float64   `json:"p_r_qty,omitempty"`
	Price                  *float64   `json:"p_r_price,omitempty"`
	TakeProfitPrice        *float64   `json:"take_profit,omitempty"`
	StopLossPrice          *float64   `json:"stop_loss,omitempty"`
	TakeProfitTriggerPrice *float64   `json:"tp_trigger_by,omitempty"`
	StopLossTriggerPrice   *float64   `json:"sl_trigger_by,omitempty"`
}

// LinearOrderReplace :
func (s *AccountService) LinearOrderReplace(param LinearOrderReplaceParam) (*LinearOrderReplaceResponse, error) {
	var res LinearOrderReplaceResponse

	url, err := s.Client.BuildPrivateURL("/private/linear/order/replace", nil)
	if err != nil {
		return nil, err
	}

	jsonBody, err := json.Marshal(param)
	if err != nil {
		return nil, fmt.Errorf("json marshal for LinearOrderReplaceParam: %w", err)
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}
	return &res, nil
}

// LinearOrderSearchResponse :
type LinearOrderSearchResponse struct {
	ListLinearPositionsResponse
}

// LinearOrderSearch (real-time) :
func (s *AccountService) LinearOrderSearch(param QueryLinearOrderParam) (*LinearOrderSearchResponse, error) {
	var res LinearOrderSearchResponse

	if param.OrderID != nil || param.OrderLinkID != nil {
		return nil, fmt.Errorf("use 'LinearOrderSearchBy' if search only one Order")
	}
	var p map[string]string
	jsonBody, err := json.Marshal(param)
	if err != nil {
		return nil, fmt.Errorf("json marshal for QueryLinearOrderParam: %w", err)
	}
	json.Unmarshal(jsonBody, &p)

	url, err := s.Client.BuildPrivateURL("/private/linear/order/search", p)
	if err != nil {
		return nil, err
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}
	return &res, nil
}

// LinearOrderSearchBy (real-time) :
// search one order by order_id or client id
func (s *AccountService) LinearOrderSearchBy(param QueryLinearOrderParam) (*CreateLinearOrderResponse, error) {
	var res CreateLinearOrderResponse

	if param.OrderID == nil && param.OrderLinkID == nil {
		return nil, fmt.Errorf("either OrderID or OrderLinkID needed")
	}
	var p map[string]string
	jsonBody, err := json.Marshal(param)
	if err != nil {
		return nil, fmt.Errorf("json marshal for QueryLinearOrderParam: %w", err)
	}
	json.Unmarshal(jsonBody, &p)

	url, err := s.Client.BuildPrivateURL("/private/linear/order/search", p)
	if err != nil {
		return nil, err
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}
	return &res, nil
}
