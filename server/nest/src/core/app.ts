import { Pager } from '../decorators/pager.decorator';

export const ResponseList = (
  data: Record<any, any>[],
  pager: Pager,
  total: number,
) => {
  return {
    list: data,
    page: pager.page,
    page_size: pager.page_size,
    total,
  };
};
