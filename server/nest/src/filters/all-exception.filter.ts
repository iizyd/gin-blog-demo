import { ArgumentsHost, Catch, ExceptionFilter, Inject } from '@nestjs/common';
import { Response, Request } from 'express';
import { WINSTON_MODULE_PROVIDER } from 'nest-winston';
import { Logger } from 'winston';

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
    const status = 500;
    const message = '服务器端错误';

    const response_object = {
      code: status,
      msg: message,
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

    response.status(500).json(response_object);
  }
}
