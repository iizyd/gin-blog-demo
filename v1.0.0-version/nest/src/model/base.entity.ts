import { Column, UpdateDateColumn } from 'typeorm';

export abstract class BaseEntity {
  @Column('tinyint')
  state: number;

  @Column('tinyint', { default: 0 })
  is_del: number;

  @Column({ type: 'datetime', default: () => 'CURRENT_TIMESTAMP' })
  created_on: Date;

  @Column()
  created_by?: string;

  @Column({ type: 'datetime', default: () => 'CURRENT_TIMESTAMP' })
  @UpdateDateColumn()
  modified_on: Date;

  @Column()
  modified_by?: string;

  @Column({ type: 'datetime' })
  deleted_on: Date;
}
