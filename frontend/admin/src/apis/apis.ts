import { Get, Post } from "./axios";

const apis = {
  getArticles: (data: any = {}) => Get("/api/v1/articles", data),

  uploadFile: (data: any = {}) =>
    Post("/upload/file", data, {
      headers: {
        "Content-Type": "multipart/form-data",
      }
    }),
};

export default apis;
