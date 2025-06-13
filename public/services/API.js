export const API = {
    baseURL: "/api",
    getTopMovies: async () => {
        return API.fetch("movies/top");
    },
    fetch: async (Service, args) => {
        try {
            const response = await fetch(API.baseURL + Service);
            const result = await response.json();
            return result;
        } catch (e) {
            console.error(e);
        }
    }
}