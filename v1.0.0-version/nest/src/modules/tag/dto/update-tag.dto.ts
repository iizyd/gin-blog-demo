import { PartialType } from '@nestjs/mapped-types';
import { CreateTagDto } from './create-tag.dto';
import { IsIn, IsOptional, Length } from 'class-validator';
import { Type } from 'class-transformer';

export class UpdateTagDto extends PartialType(CreateTagDto) {
  @Length(0, 100)
  name: string;

  @IsIn([0, 1])
  @IsOptional()
  @Type(() => Number)
  state = 0;
}
