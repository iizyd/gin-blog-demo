import { Injectable } from '@nestjs/common';
import { CreateTagDto } from './dto/create-tag.dto';
import { UpdateTagDto } from './dto/update-tag.dto';
import { InjectRepository } from '@nestjs/typeorm';
import { Tag } from './entities/tag.entity';
import { FindManyOptions, Repository } from 'typeorm';
import { ListTagDto } from './dto/list-tag.dto';
import { Pager } from '../../decorator/pager.decorator';

@Injectable()
export class TagService {
  constructor(@InjectRepository(Tag) private tagRepository: Repository<Tag>) {}

  create(createTagDto: CreateTagDto) {
    return 'This action adds a new tag';
  }

  async findAll(listTagDto: ListTagDto, pager: Pager) {
    const findOption: FindManyOptions = { where: { ...listTagDto } };
    if (pager.ok) {
      findOption.skip = pager.page_offset;
      findOption.take = pager.page_size;
    }
    console.log(findOption);
    return await this.tagRepository.find(findOption);
  }

  findOne(id: number) {
    return `This action returns a #${id} tag`;
  }

  update(id: number, updateTagDto: UpdateTagDto) {
    return `This action updates a #${id} tag`;
  }

  remove(id: number) {
    return `This action removes a #${id} tag`;
  }
}
