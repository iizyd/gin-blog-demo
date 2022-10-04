import { IsIn, IsOptional } from 'class-validator';
import { Type } from 'class-transformer';

export class BaseDto {
  @IsIn([0, 1])
  @IsOptional({})
  @Type(() => Number)
  is_del = 0;

  @IsOptional()
  created_by = '';

  @IsOptional()
  modified_by = '';
}
