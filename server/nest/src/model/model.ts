import { Column } from 'typeorm';

export abstract class Model {
  @Column('tinyint')
  state: number;

  @Column('tinyint', { default: 0 })
  is_del: number;

  @Column({ type: 'datetime', default: () => 'CURRENT_TIMESTAMP' })
  created_on: Date;

  @Column()
  created_by: string;

  @Column({ type: 'datetime', default: () => 'CURRENT_TIMESTAMP' })
  modified_on: Date;

  @Column()
  modified_by: string;

  @Column({ type: 'datetime' })
  deleted_on: Date;
}
