// Libraries
import {PureComponent} from 'react'
import _ from 'lodash'

// Types
import {FluxToolbarFunction} from 'src/types/flux'

interface Props {
  funcs: FluxToolbarFunction[]
  children: (
    sortedFunctions: {[category: string]: FluxToolbarFunction[]}
  ) => JSX.Element | JSX.Element[]
}

class TransformToolbarFunctions extends PureComponent<Props> {
  public render() {
    return this.props.children(this.sortedFunctions)
  }

  private get sortedFunctions() {
    const {funcs} = this.props
    return _.groupBy(funcs, 'category')
  }
}

export default TransformToolbarFunctions
