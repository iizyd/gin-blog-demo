import {
  CallHandler,
  ExecutionContext,
  HttpStatus,
  Injectable,
  NestInterceptor,
} from '@nestjs/common';
import { map, Observable } from 'rxjs';
import { errorCodes } from '../common/errors/code';

interface Response<T> {
  data: T;
}

@Injectable()
export class ResponseInterceptorInterceptor<T>
  implements NestInterceptor<T, Response<T>>
{
  intercept(context: ExecutionContext, next: CallHandler<T>): Observable<any> {
    return next.handle().pipe(
      map((data) => {
        return {
          status: HttpStatus.OK,
          code: errorCodes.Success.getCode(),
          msg: errorCodes.Success.getMsg(),
          data: data,
        };
      }),
    );
  }
}
