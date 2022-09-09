import { createParamDecorator, ExecutionContext } from '@nestjs/common';
import Config from '../config/config';

export interface Pager {
  page: number;
  page_size: number;
  page_offset: number;
  ok: boolean;
}

const config = Config();
const default_page = 1;
const default_page_size = config.App.DefaultPageSize;
const max_page_size = config.App.MaxPageSize;

export const Pager = createParamDecorator<never, ExecutionContext, Pager>(
  (data: never, ctx: ExecutionContext): Pager => {
    const request = ctx.switchToHttp().getRequest();
    let page = +request.query.page || default_page;
    let page_size = +request.query.page_size || default_page_size;
    let page_offset = 0;

    if (page <= 0) page = default_page;
    if (page_size <= 0) page_size = default_page_size;
    if (page_size > max_page_size) page_size = max_page_size;

    if (page > 0) page_offset = (page - 1) * page_size;

    return {
      page,
      page_size,
      page_offset,
      ok: page > 0 && page_size > 0,
    };
  },
);
