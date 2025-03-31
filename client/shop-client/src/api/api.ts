import axios from "axios";

const API_BASE_URL = "http://localhost:8080"

export const api = axios.create({
    baseURL: API_BASE_URL,
    headers: {
        "Accept": "application/json",
        "Content-Type": "application/json"
    },
});