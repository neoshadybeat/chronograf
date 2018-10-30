// Libraries
import React, {PureComponent} from 'react'

// Types
import {FluxToolbarFunction} from 'src/types/flux'

interface Props {
  func: FluxToolbarFunction
}

class FunctionsToolbarTooltip extends PureComponent<Props> {
  public render() {
    return (
      <div className="flux-functions-toolbar--tooltip">
        <div>ahoy yoooooo</div>
        {/* {this.description}
        {this.arguments}
        {this.example} */}
      </div>
    )
  }

  private get description(): JSX.Element {
    const {
      func: {desc},
    } = this.props
    return <div>{desc}</div>
  }

  private get arguments(): JSX.Element {
    return (
      <div>
        Arguments<div>args</div>
      </div>
    )
  }

  private get example(): JSX.Element {
    const {
      func: {example},
    } = this.props
    return (
      <div>
        Example<div>{example}</div>
      </div>
    )
  }
}

export default FunctionsToolbarTooltip
