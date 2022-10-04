import {
  Body,
  Controller,
  Delete,
  Get,
  HttpException,
  HttpStatus,
  Param,
  Post,
  Put,
  Query,
} from '@nestjs/common';
import { TagService } from './tag.service';
import { CreateTagDto } from './dto/create-tag.dto';
import { UpdateTagDto } from './dto/update-tag.dto';
import { Pager } from '../../decorators/pager.decorator';
import { ListTagDto } from './dto/list-tag.dto';
import { ResponseList } from '../../core/app';

@Controller('tags')
export class TagController {
  constructor(private readonly tagService: TagService) {}

  @Post()
  async create(@Body() createTagDto: CreateTagDto) {
    if (await this.tagService.create(createTagDto)) {
      return null;
    }

    throw new HttpException(
      'create tag error',
      HttpStatus.INTERNAL_SERVER_ERROR,
    );
  }

  @Get()
  async findAll(@Query() listTagDto: ListTagDto, @Pager() pager: Pager) {
    const list = await this.tagService.findAll(listTagDto, pager);
    const total = await this.tagService.count(listTagDto);

    return ResponseList(list, pager, total);
  }

  @Get(':id')
  findOne(@Param('id') id: string) {
    return this.tagService.findOne(+id);
  }

  @Put(':id')
  async update(@Param('id') id: string, @Body() updateTagDto: UpdateTagDto) {
    const res = await this.tagService.update(+id, updateTagDto);
    if (res.affected) {
      return null;
    }

    throw new HttpException(
      'update tag error',
      HttpStatus.INTERNAL_SERVER_ERROR,
    );
  }

  @Delete(':id')
  async remove(@Param('id') id: string) {
    const res = await this.tagService.remove(+id);
    if (res.affected) {
      return null;
    }

    throw new HttpException(
      'delete tag error',
      HttpStatus.INTERNAL_SERVER_ERROR,
    );
  }
}
