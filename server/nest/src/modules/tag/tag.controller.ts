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

@Controller('tag')
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
  update(@Param('id') id: string, @Body() updateTagDto: UpdateTagDto) {
    return this.tagService.update(+id, updateTagDto);
  }

  @Delete(':id')
  remove(@Param('id') id: string) {
    return this.tagService.remove(+id);
  }
}
