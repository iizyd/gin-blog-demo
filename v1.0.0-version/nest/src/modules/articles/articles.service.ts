import { Injectable } from '@nestjs/common';
import { CreateArticleDto } from './dto/create-article.dto';
import { UpdateArticleDto } from './dto/update-article.dto';
import { FindManyOptions, Repository } from 'typeorm';
import { ListArticleDto } from './dto/list-article.dto';
import { Pager } from '../../decorators/pager.decorator';
import { InjectRepository } from '@nestjs/typeorm';
import { Article } from './entities/article.entity';

@Injectable()
export class ArticlesService {
  constructor(
    @InjectRepository(Article) private articleRepository: Repository<Article>,
  ) {}

  create(createArticleDto: CreateArticleDto) {
    return 'This action adds a new article';
  }

  async findAll(listArticleDto: ListArticleDto, pager: Pager) {
    const findOption: FindManyOptions = { where: { ...listArticleDto } };
    if (pager.ok) {
      findOption.skip = pager.page_offset;
      findOption.take = pager.page_size;
    }
    return await this.articleRepository.find(findOption);
  }

  async count(listArticleDto: ListArticleDto) {
    const findOption: FindManyOptions = { where: { ...listArticleDto } };
    return await this.articleRepository.count(findOption);
  }

  findOne(id: number) {
    return `This action returns a #${id} article`;
  }

  update(id: number, updateArticleDto: UpdateArticleDto) {
    return `This action updates a #${id} article`;
  }

  remove(id: number) {
    return `This action removes a #${id} article`;
  }
}
