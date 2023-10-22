import { IsIn, IsOptional, Length } from 'class-validator';
import { Type } from 'class-transformer';
import { BaseDto } from '../../../dto/base.dto';

export class CreateTagDto extends BaseDto {
  @Length(0, 100)
  name: string;

  @IsIn([0, 1])
  @IsOptional()
  @Type(() => Number)
  state = 0;

  @IsOptional()
  created_by = '';

  @IsOptional()
  modified_by = '';
}
