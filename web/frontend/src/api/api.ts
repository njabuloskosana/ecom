import axios from "axios";
import type { AxiosInstance, AxiosRequestConfig, AxiosResponse, InternalAxiosRequestConfig } from "axios";
import Config from "@/config";
// Define the Axios instance
const apiClient: AxiosInstance = axios.create({
  baseURL:  Config.BASE_API_URL_PRODUCTION, // Use environment variable or default
  timeout: 10000, // Timeout for API requests
  headers: {
    "Content-Type": "application/json",
  },
});

// Add a request interceptor to handle tokens or other global headers
apiClient.interceptors.request.use(
  (config: InternalAxiosRequestConfig) => {
    const jwtToken = localStorage.getItem("jwtToken"); // Retrieve JWT token from storage
    if (jwtToken) {
      config.headers!.Authorization = `Bearer ${jwtToken}`; // Attach token to requests
    }
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

// Add a response interceptor to handle errors globally
apiClient.interceptors.response.use(
  (response: AxiosResponse) => {
    return response.data; // Return only the data property by default
  },
  (error) => {
    if (error.response) {
      // Handle HTTP errors (e.g., 401, 403, 500)
      console.error("API Error:", error.response.data, error.response.status);
      throw new Error(error.response.data.message || "An error occurred.");
    } else if (error.request) {
      // Handle cases where no response was received
      console.error("No response from server:", error.request);
      throw new Error("No response from the server.");
    } else {
      // Handle other errors
      console.error("Error:", error.message);
      throw new Error("An unexpected error occurred.");
    }
  }
);

export const get = <T>(url: string, config?: AxiosRequestConfig): Promise<T> =>
  apiClient.get<T>(url, config).then(response => response.data);

export const post = <T>(
  url: string,
  data?: any,
  config?: AxiosRequestConfig
): Promise<T> => apiClient.post<T>(url, data, config).then(response => response.data);

export const put = <T>(
  url: string,
  data?: any,
  config?: AxiosRequestConfig
): Promise<T> => apiClient.put<T>(url, data, config).then(response => response.data);

export const del = <T>(
  url: string,
  config?: AxiosRequestConfig
): Promise<T> => apiClient.delete<T>(url, config).then(response => response.data);