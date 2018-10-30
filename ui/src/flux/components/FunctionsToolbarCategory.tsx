// Libraries
import React, {PureComponent} from 'react'

// Components
import FunctionsToolbarFunction from 'src/flux/components/FunctionsToolbarFunction'

// Types
import {FluxToolbarFunction} from 'src/types/flux'

interface Props {
  category: string
  funcs: FluxToolbarFunction[]
}

class FunctionsToolbarCategory extends PureComponent<Props> {
  public render() {
    const {category, funcs} = this.props
    return (
      <dl className="flux-functions-toolbar--category">
        <dt>{category}</dt>
        {funcs.map(func => (
          <FunctionsToolbarFunction key={func.name} func={func} />
        ))}
      </dl>
    )
  }
}

export default FunctionsToolbarCategory
