import { HttpException, HttpStatus } from '@nestjs/common';
import { CustomError } from './custom-error';

export class CustomException extends HttpException {
  constructor(codeEntity: CustomError, httpStatus: HttpStatus) {
    super(codeEntity, httpStatus);
  }
}
