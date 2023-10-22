import {
  ArgumentsHost,
  Catch,
  ExceptionFilter,
  HttpException,
  Inject,
} from '@nestjs/common';
import { Response, Request } from 'express';
import { WINSTON_MODULE_PROVIDER } from 'nest-winston';
import { Logger } from 'winston';

/**
 * 捕获 http 异常
 */
@Catch(HttpException)
export class HttpExceptionFilter implements ExceptionFilter<HttpException> {
  constructor(
    @Inject(WINSTON_MODULE_PROVIDER)
    private readonly logger: Logger,
  ) {}

  catch(exception: HttpException, host: ArgumentsHost) {
    const ctx = host.switchToHttp();
    const response = ctx.getResponse<Response>();
    const request = ctx.getRequest<Request>();
    const status = exception.getStatus();
    const exception_response = exception.getResponse() as any;
    const message =
      exception_response.message || exception_response.msg || exception.message;

    const response_object = {
      status: status,
      code: exception_response.code || -1,
      msg: message,
      data: null,
    };

    this.logger.error({
      message: exception.stack,
      meta: [
        { ip: request.ip },
        { request_path: request.originalUrl },
        { method: request.method },
        { query: request.query },
        { params: request.params },
        { body: request.body },
        { response_body: response_object },
      ],
    });

    response.status(status).json(response_object);
  }
}
