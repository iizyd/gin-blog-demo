import { createParamDecorator, ExecutionContext } from '@nestjs/common';

export interface Pager {
  page: number;
  page_size: number;
  page_offset: number;
}
export const Pager = createParamDecorator<never, ExecutionContext, Pager>(
  (data: never, ctx: ExecutionContext): Pager => {
    const request = ctx.switchToHttp().getRequest();
    let page = request.query.page;
    let page_size = request.query.page_size;
    let page_offset = 0;

    if (page <= 0) page = 1;
    if (page_size <= 0) page_size = 10;
    if (page_size > 100) page_size = 100;

    if (page > 0) page_offset = (page - 1) * page_size;

    return {
      page,
      page_size,
      page_offset,
    };
  },
);
