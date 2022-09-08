import {
  ArgumentsHost,
  Catch,
  ExceptionFilter,
  HttpStatus,
  Inject,
} from '@nestjs/common';
import { Response, Request } from 'express';
import { WINSTON_MODULE_PROVIDER } from 'nest-winston';
import { Logger } from 'winston';
import { errorCodes } from '../common/errors/code';

/**
 * 捕获 非http 异常
 * 在 HttpExceptionFilter 之前执行
 * 如果发生 非http 异常，则直接返回 500
 */
@Catch()
export class AllExceptionFilter implements ExceptionFilter<Error> {
  constructor(
    @Inject(WINSTON_MODULE_PROVIDER)
    private readonly logger: Logger,
  ) {}

  catch(exception: Error, host: ArgumentsHost) {
    const ctx = host.switchToHttp();
    const response = ctx.getResponse<Response>();
    const request = ctx.getRequest<Request>();
    const status = HttpStatus.INTERNAL_SERVER_ERROR;

    const response_object = {
      status,
      code: errorCodes.ServerError.getCode(),
      msg: errorCodes.ServerError.getMsg(),
      data: null,
    };

    // 向控制台输出
    console.error(exception);

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
