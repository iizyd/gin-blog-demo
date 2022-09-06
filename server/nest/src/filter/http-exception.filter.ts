import {
  ArgumentsHost,
  Catch,
  ExceptionFilter,
  HttpException,
} from '@nestjs/common';
import { Response } from 'express';

@Catch(HttpException)
export class HttpExceptionFilter implements ExceptionFilter<HttpException> {
  catch(exception: HttpException, host: ArgumentsHost) {
    const ctx = host.switchToHttp();
    const response = ctx.getResponse<Response>();
    const status = exception.getStatus();
    const exception_response = exception.getResponse() as any;
    const message = exception_response.message
      ? exception_response.message
      : exception.message;

    response.status(status).json({
      code: status,
      msg: message,
      data: null,
    });
  }
}
