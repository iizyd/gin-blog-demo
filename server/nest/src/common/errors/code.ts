import { createError } from './custom-error';
import { ErrorCodes } from './types';

export const errorCodes: ErrorCodes = Object.freeze({
  Success: createError(0, '成功'),
  ServerError: createError(10000000, '服务器内部错误'),
  InvalidParams: createError(10000001, '入参错误'),
  NotFound: createError(10000002, '找不到'),
  UnauthorizedAuthNotExit: createError(10000003, '鉴权失败，用户名或密码错误'),
  UnauthorizedTokenError: createError(10000004, '鉴权失败，Token 错误'),
  UnauthorizedTokenTimeout: createError(10000005, '鉴权失败，Token 过期'),
  UnauthorizedTokenGenerate: createError(10000006, '鉴权失败，Token 生成失败'),
  TooManyRequest: createError(10000007, '请求过多'),
});
