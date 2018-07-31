import React, {Component} from 'react'
import _ from 'lodash'

import Dropdown from 'src/shared/components/Dropdown'

import {getDeep} from 'src/utils/wrappers'

import {Source, Cell, CellQuery} from 'src/types'
import {ImportedSources} from 'src/types/dashboards'

interface Props {
  cells: Cell[]
  source: Source
  sources: Source[]
  importedSources: ImportedSources
  onSubmit: (cells: Cell[]) => void
}

interface State {
  sourcesCells: SourcesCells
  sourceMappings: SourceMappings
}

interface CellInfo {
  id: string
  name: string
}

interface SourcesCells {
  [x: string]: CellInfo[]
}

interface SourceInfo {
  name: string
  id: string
  link: string
}

interface SourceMappings {
  [x: string]: SourceInfo
}

interface SourceItemValue {
  importedSourceID: string
  sourceInfo: SourceInfo
  text?: string
}

const NO_SOURCE = 'none'

class ImportDashboardMappings extends Component<Props, State> {
  constructor(props: Props) {
    super(props)
    this.state = {sourcesCells: {}, sourceMappings: {}}
  }

  public componentDidMount() {
    const {cells, importedSources, source} = this.props

    if (_.isEmpty(cells)) {
      return
    }

    let sourcesCells: SourcesCells = {}
    const sourceMappings: SourceMappings = {}
    const sourceInfo: SourceInfo = this.getSourceInfo(source)
    const cellsWithNoSource: CellInfo[] = []

    sourcesCells = _.reduce(
      cells,
      (acc, c) => {
        const cellInfo: CellInfo = {id: c.i, name: c.name}
        const query = getDeep<CellQuery>(c, 'queries.0', null)
        if (_.isEmpty(query)) {
          return acc
        }

        const sourceLink = getDeep<string>(query, 'source', '')
        if (!sourceLink) {
          cellsWithNoSource.push(cellInfo)
          return acc
        }

        let importedSourceID = _.findKey(
          importedSources,
          is => is.link === sourceLink
        )
        if (!importedSourceID) {
          const sourceIDRegex = /sources\/(\d+)/g
          // first capture group
          const sourceLinkSID = sourceIDRegex.exec(sourceLink)[1]
          if (!sourceLinkSID) {
            return acc
          }
          importedSourceID = sourceLinkSID
        }

        if (acc[importedSourceID]) {
          acc[importedSourceID].push(cellInfo)
        } else {
          acc[importedSourceID] = [cellInfo]
        }
        sourceMappings[importedSourceID] = sourceInfo
        return acc
      },
      sourcesCells
    )

    if (cellsWithNoSource.length) {
      sourcesCells[NO_SOURCE] = cellsWithNoSource
    }

    this.setState({sourcesCells, sourceMappings})
  }

  public render() {
    return (
      <div className="container-fluid">
        <div className="row">
          <div className="col-md-12">
            <div className="panel">
              <div className="panel-body">
                {this.table}
                <button
                  className="btn btn-sm btn-success"
                  onClick={this.handleSubmit}
                >
                  Done
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    )
  }

  private get noMappings(): JSX.Element {
    return <div data-test="no-mapping">No source mappings required</div>
  }

  private get table(): JSX.Element {
    const {sourcesCells} = this.state

    if (_.isEmpty(sourcesCells)) {
      return this.noMappings
    }

    return (
      <table className="table v-center admin-table">
        {this.header}
        <tbody>{this.tableBody}</tbody>
      </table>
    )
  }

  private get tableBody(): JSX.Element[] {
    const {importedSources} = this.props
    const {sourcesCells} = this.state

    const rows = _.reduce(
      sourcesCells,
      (acc, __, i) => {
        if (i !== NO_SOURCE && sourcesCells[i]) {
          const sourceName = getDeep<string>(
            importedSources,
            `${i}.name`,
            'Source'
          )
          acc.push(this.getRow(sourceName, i))
        }
        return acc
      },
      []
    )
    if (sourcesCells[NO_SOURCE]) {
      const noSourceRow = this.getRow('No Source Found', NO_SOURCE)
      rows.push(noSourceRow)
    }
    return rows
  }

  private getRow(sourceName: string, sourceID: string): JSX.Element {
    let sourceLabel = `${sourceName} (${sourceID})`
    let description = 'Cells that use this Source:'
    if (sourceID === NO_SOURCE) {
      sourceLabel = sourceName
      description = 'Cells with no Source:'
    }
    return (
      <tr key={sourceID}>
        <td>
          <div data-test="source-label">{sourceLabel}</div>
          <div>{description}</div>
          {this.getCellsForSource(sourceID)}
        </td>
        <td>
          <div className="fancytable--td provider--arrow">
            <span />
          </div>
        </td>
        <td className="text-right">
          {/* <Dropdown
            onChange={this.handleDropdownChange}
            selectedID={this.getSelectedSourceID(sourceID)}
          >
            {this.getSourceItems(sourceID)}
          </Dropdown> */}
          <Dropdown
            className="dropdown-stretch"
            buttonColor="btn-default"
            buttonSize="btn-xs"
            items={this.getSourceItems(sourceID)}
            onChoose={this.handleDropdownChange}
            selected={this.getSelected(sourceID)}
          />
        </td>
      </tr>
    )
  }

  private getSourceItems(importedSourceID: string) {
    const {sources} = this.props

    return sources.map(source => {
      const sourceInfo = this.getSourceInfo(source)
      const sourceMap: SourceItemValue = {
        sourceInfo,
        importedSourceID,
        text: source.name,
      }
      return sourceMap
    })
  }

  private getSourceInfo(source: Source): SourceInfo {
    return {
      name: source.name,
      id: source.id,
      link: source.links.self,
    }
  }

  private get header() {
    return (
      <thead>
        <tr>
          <th>Sources in Dashboard</th>
          <th />
          <th>Available Sources</th>
        </tr>
      </thead>
    )
  }

  private getSelected(importedSourceID: string): string {
    const {sources} = this.props
    const {sourceMappings} = this.state

    const sourceMapping = sourceMappings[importedSourceID]
    if (sourceMapping) {
      return sourceMappings[importedSourceID].name
    }

    return sources[0].name
  }

  private getCellsForSource(sourceID): JSX.Element[] {
    const {sourcesCells} = this.state

    return _.map(sourcesCells[sourceID], c => {
      return <div key={c.id}>{c.name}</div>
    })
  }

  private handleDropdownChange = (item: SourceItemValue): void => {
    const {sourceMappings} = this.state

    sourceMappings[item.importedSourceID] = item.sourceInfo
    this.setState({sourceMappings})
  }

  private handleSubmit = (): void => {
    const {cells, onSubmit, importedSources} = this.props
    const {sourceMappings} = this.state

    const mappedCells = cells.map(c => {
      const query = getDeep<CellQuery>(c, 'queries.0', null)
      if (_.isEmpty(query)) {
        return c
      }
      const sourceLink = getDeep<string>(query, 'source', '')
      if (!sourceLink) {
        const mappedSourceLink = sourceMappings[NO_SOURCE].link
        let queries = getDeep<CellQuery[]>(c, 'queries', [])
        if (queries.length) {
          queries = queries.map(q => {
            return {...q, source: mappedSourceLink}
          })
        }
        return {...c, queries}
      }
      const importedSourceID = _.findKey(
        importedSources,
        is => is.link === sourceLink
      )
      if (sourceLink && importedSourceID) {
        const mappedSourceLink = sourceMappings[importedSourceID].link
        let queries = getDeep<CellQuery[]>(c, 'queries', [])
        if (queries.length) {
          queries = queries.map(q => {
            return {...q, source: mappedSourceLink}
          })
        }
        return {...c, queries}
      }

      return c
    })

    onSubmit(mappedCells)
  }
}

export default ImportDashboardMappings