export interface Ticker {
    ticker: string;
    company: string;
    price: number;
    logo: string;
    market_cap: number;
    last_review: string;
    sentiment: string;
}

export interface Meta {
    page: number;
    total: number;
    last_page: number;
    per_page: number;
}
