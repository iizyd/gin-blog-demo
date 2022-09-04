import { IsIn, IsNumber, IsOptional, Length } from 'class-validator';
import { Type } from 'class-transformer';

export class ListTagDto {
  @Length(0, 100)
  @IsOptional()
  name: string;

  @IsIn([0, 1])
  @IsOptional()
  @Type(() => Number)
  state: number;
}
