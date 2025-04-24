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

export interface Recommendation {
    time: string;
    action: string;
    rating_from: string;
    rating_from_txt: string;
    rating_to: string;
    rating_to_txt: string;
    target_from: number;
    target_to: number;
    broker: string;
}
