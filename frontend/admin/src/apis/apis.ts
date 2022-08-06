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

  // 获取标签列表
  getTags: (data: any = {}) => Get("/api/v1/tags", data),

  // 创建标签
  createTag: (data: { name: string; state?: number; created_by: string }) =>
    Post(`/api/v1/tags`, data),

  // 修改标签
  updateTag: (
    id: number,
    data: { name?: string; state?: number; modified_by: string }
  ) => Put(`/api/v1/tags/${id}`, data),

  // 删除标签
  delTag: (id: string | number) => Del(`/api/v1/tags/${id}`, {}),

  // 登录
  Login: (data: { user_name: string; password: string }) => Post("/auth", data),
};

export default apis;
