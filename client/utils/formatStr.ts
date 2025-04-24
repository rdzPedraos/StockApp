export const formatPrice = (price: number) => {
    console.log(price);
    if (Math.abs(price) >= 1000000) {
        const suffixes = ['', 'K', 'M', 'B', 'T'];
        const suffixNum = Math.floor(Math.log10(Math.abs(price)) / 3);
        const shortValue = (price / Math.pow(1000, suffixNum)).toFixed(1);

        return `$${shortValue} ${suffixes[suffixNum]}`;
    }

    // Para valores normales, usamos el formateador estándar
    return price.toLocaleString('en-US', {
        style: 'currency',
        currency: 'USD',
    });
};

export const formatDate = (date: string) => {
    return new Date(date).toLocaleDateString('en-US', {
        month: 'long',
        day: 'numeric',
        year: 'numeric',
    });
};

export const formatPercentage = (percentage: number) => {
    return percentage.toFixed(2) + '%';
};
