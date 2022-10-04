import { Pager } from '../decorators/pager.decorator';
import dayjs from 'dayjs';

export const ResponseList = (
  data: Record<any, any>[],
  pager: Pager,
  total: number,
) => {
  if (Array.isArray(data)) {
    data.forEach((item) => {
      item.created_on && (item.created_on = timeTransform(item.created_on));
      item.modified_on && (item.modified_on = timeTransform(item.modified_on));
      item.deleted_on && (item.deleted_on = timeTransform(item.deleted_on));
    });
  }

  return {
    list: data,
    page: pager.page,
    page_size: pager.page_size,
    total,
  };
};

export const timeTransform = (timeStr: string): string => {
  return dayjs(timeStr).format('YYYY-MM-DD HH:mm:ss');
};
