import { CustomError } from './custom-error';

export interface ErrorCodes {
  Success: CustomError;
  ServerError: CustomError;
  InvalidParams: CustomError;
  NotFound: CustomError;
  UnauthorizedAuthNotExit: CustomError;
  UnauthorizedTokenError: CustomError;
  UnauthorizedTokenTimeout: CustomError;
  UnauthorizedTokenGenerate: CustomError;
  TooManyRequest: CustomError;
}
