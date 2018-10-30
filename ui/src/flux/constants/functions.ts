// Types
import {FluxToolbarFunction} from 'src/types/flux'

export const functions: FluxToolbarFunction[] = [
  {
    name: 'count()',
    args: [
      {
        name: 'columns',
        desc:
          'A list of columns on which to operate. Defaults to `["_value"]`.',
        type: 'Array of Strings',
      },
    ],
    desc:
      'The `count()` function outputs the number of non-null records in each aggregated column.',
    example: 'count(columns: ["_value"])',
    category: 'Aggregate',
    link: 'http://example.com',
  },
  {
    name: 'covariance()',
    args: [
      {
        name: 'columns',
        desc:
          'A list of columns on which to operate. Exactly two columns must be provided.',
        type: 'Array of Strings',
      },
      {
        name: 'pearsonr',
        desc:
          'Indicates whether the result should be normalized to be the Pearson R coefficient',
        type: 'Boolean',
      },
      {
        name: 'valueDst',
        desc:
          'The column into which the result will be placed. Defaults to `"_value"`.',
        type: 'String',
      },
    ],
    desc:
      'The `covariance()` function computes the covariance between two columns.',
    example:
      'covariance(columns: ["column_x", "column_y"], pearsonr: false, valueDst: "_value")',
    category: 'Aggregate',
    link: 'http://example.com',
  },
  {
    name: 'from()',
    args: [
      {
        name: 'columns',
        desc:
          'A list of columns on which to operate. Defaults to `["_value"]`.',
        type: 'Array of Strings',
      },
    ],
    desc:
      'The `count()` function outputs the number of non-null records in each aggregated column.',
    example: 'count(columns: ["_value"])',
    category: 'Data',
    link: 'http://example.com',
  },
  {
    name: 'yield()',
    args: [
      {
        name: 'columns',
        desc:
          'A list of columns on which to operate. Defaults to `["_value"]`.',
        type: 'Array of Strings',
      },
    ],
    desc:
      'The `count()` function outputs the number of non-null records in each aggregated column.',
    example: 'count(columns: ["_value"])',
    category: 'Data',
    link: 'http://example.com',
  },
  {
    name: 'filter()',
    args: [
      {
        name: 'columns',
        desc:
          'A list of columns on which to operate. Defaults to `["_value"]`.',
        type: 'Array of Strings',
      },
    ],
    desc:
      'The `count()` function outputs the number of non-null records in each aggregated column.',
    example: 'count(columns: ["_value"])',
    category: 'Windows',
    link: 'http://example.com',
  },
  {
    name: 'range()',
    args: [
      {
        name: 'columns',
        desc:
          'A list of columns on which to operate. Defaults to `["_value"]`.',
        type: 'Array of Strings',
      },
    ],
    desc:
      'The `count()` function outputs the number of non-null records in each aggregated column.',
    example: 'count(columns: ["_value"])',
    category: 'Windows',
    link: 'http://example.com',
  },
  {
    name: 'window()',
    args: [
      {
        name: 'columns',
        desc:
          'A list of columns on which to operate. Defaults to `["_value"]`.',
        type: 'Array of Strings',
      },
    ],
    desc:
      'The `count()` function outputs the number of non-null records in each aggregated column.',
    example: 'count(columns: ["_value"])',
    category: 'Windows',
    link: 'http://example.com',
  },
  {
    name: 'count()',
    args: [
      {
        name: 'columns',
        desc:
          'A list of columns on which to operate. Defaults to `["_value"]`.',
        type: 'Array of Strings',
      },
    ],
    desc:
      'The `count()` function outputs the number of non-null records in each aggregated column.',
    example: 'count(columns: ["_value"])',
    category: 'Aggregation',
    link: 'http://example.com',
  },
  {
    name: 'cumulativeSum()',
    args: [
      {
        name: 'columns',
        desc:
          'A list of columns on which to operate. Defaults to `["_value"]`.',
        type: 'Array of Strings',
      },
    ],
    desc:
      'The `count()` function outputs the number of non-null records in each aggregated column.',
    example: 'count(columns: ["_value"])',
    category: 'Aggregation',
    link: 'http://example.com',
  },
  {
    name: 'window()',
    args: [
      {
        name: 'columns',
        desc:
          'A list of columns on which to operate. Defaults to `["_value"]`.',
        type: 'Array of Strings',
      },
    ],
    desc:
      'The `count()` function outputs the number of non-null records in each aggregated column.',
    example: 'count(columns: ["_value"])',
    category: 'Aggregation',
    link: 'http://example.com',
  },
  {
    name: 'integral()',
    args: [
      {
        name: 'columns',
        desc:
          'A list of columns on which to operate. Defaults to `["_value"]`.',
        type: 'Array of Strings',
      },
    ],
    desc:
      'The `count()` function outputs the number of non-null records in each aggregated column.',
    example: 'count(columns: ["_value"])',
    category: 'Aggregation',
    link: 'http://example.com',
  },
  {
    name: 'min()',
    args: [
      {
        name: 'columns',
        desc:
          'A list of columns on which to operate. Defaults to `["_value"]`.',
        type: 'Array of Strings',
      },
    ],
    desc:
      'The `count()` function outputs the number of non-null records in each aggregated column.',
    example: 'count(columns: ["_value"])',
    category: 'Aggregation',
    link: 'http://example.com',
  },
  {
    name: 'mean()',
    args: [
      {
        name: 'columns',
        desc:
          'A list of columns on which to operate. Defaults to `["_value"]`.',
        type: 'Array of Strings',
      },
    ],
    desc:
      'The `count()` function outputs the number of non-null records in each aggregated column.',
    example: 'count(columns: ["_value"])',
    category: 'Aggregation',
    link: 'http://example.com',
  },
  {
    name: 'mode()',
    args: [
      {
        name: 'columns',
        desc:
          'A list of columns on which to operate. Defaults to `["_value"]`.',
        type: 'Array of Strings',
      },
    ],
    desc:
      'The `count()` function outputs the number of non-null records in each aggregated column.',
    example: 'count(columns: ["_value"])',
    category: 'Aggregation',
    link: 'http://example.com',
  },
]
