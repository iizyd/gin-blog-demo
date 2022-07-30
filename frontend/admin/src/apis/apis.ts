import { Get } from "./axios";

const apis = {
  getArticles: (data: any = {}) => Get("/api/v1/articles", data),
};

export default apis;
