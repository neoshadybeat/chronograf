// Libraries
import React, {PureComponent} from 'react'

// Component
import FunctionsToolbarTooltip from 'src/flux/components/FunctionsToolbarToolTip'

// Types
import {FluxToolbarFunction} from 'src/types/flux'

interface Props {
  func: FluxToolbarFunction
}

interface State {
  isActive: boolean
}

class FunctionsToolbarFunction extends PureComponent<Props, State> {
  constructor(props: Props) {
    super(props)

    this.state = {isActive: false}
  }
  public render() {
    const {func} = this.props
    return (
      <div className="flux-functions-toolbar--function">
        {this.tooltip}
        <dd onMouseOver={this.handleHover} onMouseOut={this.handleStopHover}>
          {func.name}
        </dd>
      </div>
    )
  }

  private get tooltip(): JSX.Element {
    if (this.state.isActive) {
      return <FunctionsToolbarTooltip func={this.props.func} />
    }
  }

  private handleHover = () => {
    this.setState({isActive: true})
  }

  private handleStopHover = () => {
    this.setState({isActive: false})
  }
}

export default FunctionsToolbarFunction
