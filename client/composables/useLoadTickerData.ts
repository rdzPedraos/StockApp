import type { Ticker } from '~/shared/types/models';

const btc = {
    ticker: 'BTC',
    company: 'bitcoin',
    price: 85242.44,
    logo: 'https://s2.coinmarketcap.com/static/img/coins/64x64/1.png',
    market_cap: 1600000000000,
    last_review: '2021-01-01',
    sentiment: 'buy',
};

const eth = {
    ticker: 'ETH',
    company: 'ethereum',
    price: 2800.0,
    logo: 'https://s2.coinmarketcap.com/static/img/coins/64x64/1027.png',
    market_cap: 1600000000000,
    last_review: '2021-01-01',
    sentiment: 'sell',
};

const xrp = {
    ticker: 'XRP',
    company: 'ripple',
    price: 1.0,
    logo: 'https://s2.coinmarketcap.com/static/img/coins/64x64/52.png',
    market_cap: 1600000000000,
    last_review: '2021-01-01',
    sentiment: 'hold',
};

const items = [
    btc,
    eth,
    btc,
    xrp,
    btc,
    eth,
    xrp,
    eth,
    btc,
    xrp,
    btc,
    eth,
    btc,
    xrp,
    btc,
    eth,
    xrp,
    eth,
    btc,
    xrp,
];

export function useLoadTickerData(): Ticker[] {
    return items;
}
