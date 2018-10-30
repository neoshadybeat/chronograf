// Libraries
import React, {PureComponent} from 'react'

// Components
import TransformToolbarFunctions from 'src/flux/components/TransformToolbarFunctions'
import FunctionsToolbarCategory from 'src/flux/components/FunctionsToolbarCategory'

// Constants
import {functions as FUNCTIONS} from 'src/flux/constants'

class FluxFunctionsToolbar extends PureComponent {
  public render() {
    return (
      <div className="flux-functions-toolbar">
        <TransformToolbarFunctions funcs={FUNCTIONS}>
          {sortedFunctions => {
            return Object.entries(sortedFunctions).map(([category, funcs]) => {
              return (
                <FunctionsToolbarCategory
                  key={category}
                  category={category}
                  funcs={funcs}
                />
              )
            })
          }}
        </TransformToolbarFunctions>
      </div>
    )
  }
}

export default FluxFunctionsToolbar
