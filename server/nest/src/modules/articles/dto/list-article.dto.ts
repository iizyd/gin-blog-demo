import { IsIn, IsOptional, Length } from 'class-validator';
import { Type } from 'class-transformer';
import { BaseDto } from '../../../dto/base.dto';

export class ListArticleDto extends BaseDto {
  @Length(0, 100)
  @IsOptional()
  title: string;

  @IsOptional()
  desc: string;

  @IsOptional()
  content: string;

  @IsIn([0, 1])
  @IsOptional()
  @Type(() => Number)
  state: number;
}
