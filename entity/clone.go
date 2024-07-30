package entity

import "time"

type AdLabel struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// BidStrategy is the strategy used for bidding in the campaign.
type BidStrategy string

const (
	LowestCostWithoutCap  BidStrategy = "MENOR_CUSTO_SEM_CAP"
	LowestCostWithBidCap  BidStrategy = "MENOR_CUSTO_COM_CAP_DE_LANCE"
	CostCap               BidStrategy = "LIMITE_DE_CUSTO"
	LowestCostWithROASMin BidStrategy = "MENOR_CUSTO_COM_ROAS_MIN"
)

// BudgetScheduleSpec represents the budget schedule specifications for high-demand periods.
type BudgetScheduleSpec struct {
	TimeStart       int64  `json:"time_start"`
	TimeEnd         int64  `json:"time_end"`
	BudgetValue     int    `json:"budget_value"`
	BudgetValueType string `json:"budget_value_type"`
}

// BuyingType represents the type of buying for the campaign.
type BuyingType string

const (
	Auction  BuyingType = "AUCTION"
	Reserved BuyingType = "RESERVED"
)

// CampaignOptimizationType is the type of optimization for the campaign.
type CampaignOptimizationType string

const (
	None    CampaignOptimizationType = "NENHUM"
	IcoOnly CampaignOptimizationType = "ICO_SOMENTE"
)

// ExecutionOption represents the execution options for the campaign.
type ExecutionOption string

const (
	ValidateOnly           ExecutionOption = "validate_only"
	IncludeRecommendations ExecutionOption = "include_recommendations"
)

// Objective represents the objective of the campaign.
type Objective string

const (
	AppInstalls         Objective = "INSTALACOES_DE_APLICATIVOS"
	BrandAwareness      Objective = "CONSCIENCIA_DA_MARCA"
	Conversions         Objective = "CONVERSOES"
	EventResponses      Objective = "RESPOSTAS_DE_EVENTOS"
	LeadGeneration      Objective = "GERACAO_DE_LEADS"
	LinkClicks          Objective = "CLIQUES_DE_LINKS"
	LocalAwareness      Objective = "CONSCIENCIA_LOCAL"
	Messages            Objective = "MENSAGENS"
	OfferClaims         Objective = "REIVINDICACOES_DE_OFERTA"
	AppPromotion        Objective = "RESULTADO_PROMOCAO_DE_APLICATIVO"
	ResultAwareness     Objective = "CONSCIENCIA_DE_RESULTADO"
	ResultEngagement    Objective = "ENGAJAMENTO_DE_RESULTADO"
	ResultLeads         Objective = "LEADS_DE_RESULTADO"
	ResultSales         Objective = "VENDAS_DE_RESULTADO"
	ResultTraffic       Objective = "TRAFICO_DE_RESULTADO"
	PageLikes           Objective = "CURTIDAS_DE_PAGINA"
	PostEngagement      Objective = "ENGAJAMENTO_DE_POS-COMPUTADOR"
	ProductCatalogSales Objective = "VENDAS_DE_CATALOGO_DE_PRODUTOS"
	Reach               Objective = "ALCANCE"
	StoreVisits         Objective = "VISITAS_A_LOJA"
	VideoViews          Objective = "VISITACOES_DE_VIDEO"
)

// SpecialAdCategory is the special ad category for the campaign.
type SpecialAdCategory string

// Campaign represents a Facebook campaign.
type CampaignClone struct {
	AdLabels                  []AdLabel                `json:"adlabels"`
	BidStrategy               BidStrategy              `json:"bid_strategy"`
	BudgetScheduleSpecs       []BudgetScheduleSpec     `json:"budget_schedule_specs"`
	BuyingType                BuyingType               `json:"buying_type"`
	CampaignOptimizationType  CampaignOptimizationType `json:"campaign_optimization_type"`
	DailyBudget               int64                    `json:"daily_budget"`
	ExecutionOptions          []ExecutionOption        `json:"execution_options"`
	IsSKAdNetworkAttribution  bool                     `json:"is_skadnetwork_attribution"`
	IsUsingL3Schedule         bool                     `json:"is_using_l3_schedule"`
	IterativeSplitTestConfigs []interface{}            `json:"iterative_split_test_configs"`
	LifetimeBudget            int64                    `json:"lifetime_budget"`
	Name                      string                   `json:"name"`
	Objective                 Objective                `json:"objective"`
	PromotedObject            interface{}              `json:"promoted_object"`
	SourceCampaignID          int64                    `json:"source_campaign_id"`
	SpecialAdCategories       []SpecialAdCategory      `json:"special_ad_categories"`
	SpecialAdCategoryCountry  []string                 `json:"special_ad_category_country"`
	SpendCap                  int64                    `json:"spend_cap"`
	StartTime                 time.Time                `json:"start_time"`
	Status                    string                   `json:"status"`
	StopTime                  time.Time                `json:"stop_time"`
	ToplineID                 int64                    `json:"topline_id"`
}
