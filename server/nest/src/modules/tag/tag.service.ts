import { Injectable } from '@nestjs/common';
import { CreateTagDto } from './dto/create-tag.dto';
import { UpdateTagDto } from './dto/update-tag.dto';
import { InjectRepository } from '@nestjs/typeorm';
import { Tag } from './entities/tag.entity';
import { FindManyOptions, Repository } from 'typeorm';
import { ListTagDto } from './dto/list-tag.dto';
import { Pager } from '../../decorators/pager.decorator';

@Injectable()
export class TagService {
  constructor(@InjectRepository(Tag) private tagRepository: Repository<Tag>) {}

  async create(createTagDto: CreateTagDto) {
    return await this.tagRepository.save(createTagDto);
  }

  async findAll(listTagDto: ListTagDto, pager: Pager) {
    const findOption: FindManyOptions = { where: { ...listTagDto } };
    if (pager.ok) {
      findOption.skip = pager.page_offset;
      findOption.take = pager.page_size;
    }
    return await this.tagRepository.find(findOption);
  }

  async count(listTagDto: ListTagDto) {
    const findOption: FindManyOptions = { where: { ...listTagDto } };
    return await this.tagRepository.count(findOption);
  }

  findOne(id: number) {
    return `This action returns a #${id} tag`;
  }

  update(id: number, updateTagDto: UpdateTagDto) {
    return this.tagRepository.update(id, updateTagDto);
  }

  remove(id: number) {
    return `This action removes a #${id} tag`;
  }
}
