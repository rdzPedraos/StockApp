/**
 * Crea una versión debounced de una función
 * @param fn Función a ejecutar después del debounce
 * @param delay Tiempo de espera en milisegundos
 * @returns Función debounced
 */
// eslint-disable-next-line @typescript-eslint/no-explicit-any
export function debounce<F extends (...args: any[]) => void>(
    fn: F,
    delay = 1000
): (...args: Parameters<F>) => void {
    let timeoutId: ReturnType<typeof setTimeout> | null = null;

    return function debouncedFn(...args: Parameters<F>) {
        if (timeoutId) {
            clearTimeout(timeoutId);
        }

        timeoutId = setTimeout(() => {
            fn(...args);
            timeoutId = null;
        }, delay);
    };
}
