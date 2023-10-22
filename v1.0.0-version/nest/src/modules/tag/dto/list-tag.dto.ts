import { IsIn, IsOptional, Length } from 'class-validator';
import { Type } from 'class-transformer';
import { BaseDto } from '../../../dto/base.dto';

export class ListTagDto extends BaseDto {
  @Length(0, 100)
  @IsOptional()
  name: string;

  @IsIn([0, 1])
  @IsOptional()
  @Type(() => Number)
  state: number;
}
