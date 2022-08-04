import { Del, Get, Post, Put } from "./axios";

const apis = {
  // 获取文章列表
  getArticles: (data: any = {}) => Get("/api/v1/articles", data),

  // 获取单个文章
  getArticle: (id: string | number) => Get(`/api/v1/articles/${id}`, {}),

  // 删除单个文章
  delArticle: (id: string | number) => Del(`/api/v1/articles/${id}`, {}),

  // 修改文章
  updateArticle: (id: string | number, data: any) =>
    Put(`/api/v1/articles/${id}`, data),

  // 创建文章
  createArticle: (data: any) => Post(`/api/v1/articles`, data),

  // 上传文件
  uploadFile: (data: any = {}) =>
    Post("/upload/file", data, {
      headers: {
        "Content-Type": "multipart/form-data",
      },
    }),
};

export default apis;
