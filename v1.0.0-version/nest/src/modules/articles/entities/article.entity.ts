import { BaseEntity } from '../../../model/base.entity';
import { Column, Entity, PrimaryGeneratedColumn } from 'typeorm';

@Entity('article')
export class Article extends BaseEntity {
  @PrimaryGeneratedColumn()
  id: number;

  @Column()
  title: string;

  @Column()
  desc: string;

  @Column()
  cover_image_url: string;

  @Column()
  content: string;
}
