import { Column, Entity, PrimaryGeneratedColumn } from 'typeorm';
import { Model } from '../../../model/model';

@Entity('tag')
export class Tag extends Model {
  @PrimaryGeneratedColumn()
  id: number;

  @Column({ length: 100 })
  name: string;
}
